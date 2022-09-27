package chart

import (
	"context"

	"golang.org/x/xerrors"

	"simon/mall/service/internal/constant"
	"simon/mall/service/internal/errs"
	"simon/mall/service/internal/model/bo"
	"simon/mall/service/internal/model/po"
	"simon/mall/service/internal/utils/ctxs"
	"simon/mall/service/internal/utils/timelogger"
)

type IMemberChartUseCase interface {
	GetMemberChart(ctx context.Context) ([]*bo.MemberChart, error)
	UpdateMemberChart(ctx context.Context, cond *bo.MemberChartUpdateCond) error
	CreateMemberChart(ctx context.Context, cond *bo.MemberChartCreateCond) error
	DeleteMemberChart(ctx context.Context, cond *bo.MemberChartDelCond) error
}

type memberChartUseCase struct {
	in digIn
}

func newMemberChartUseCase(in digIn) IMemberChartUseCase {
	return &memberChartUseCase{in: in}
}

func (uc *memberChartUseCase) GetMemberChart(ctx context.Context) ([]*bo.MemberChart, error) {
	defer timelogger.LogTime(ctx)()

	memberInfo, ok := ctxs.GetSession(ctx)
	if !ok {
		return nil, errs.MemberTokenError
	}

	db := uc.in.DB.Session()
	charts, err := uc.in.MemberChartRepo.GetList(ctx, db, &po.MemberChartSearch{MemberId: memberInfo.Id})
	if err != nil {
		return nil, xerrors.Errorf("memberChartUseCase.GetMemberChart -> MemberChartRepo.GetList : %w", err)
	}

	products, err := uc.in.ProductCommon.GetProduct(ctx)
	if err != nil {
		return nil, xerrors.Errorf("memberChartUseCase.GetMemberChart -> ProductCommon.GetProduct : %w", err)
	}

	result := make([]*bo.MemberChart, len(charts))
	// use chart & product combination result
	for i := 0; i < len(charts); i++ {
		result[i] = &bo.MemberChart{
			Id:       charts[i].Id,
			Quantity: charts[i].Quantity,
		}

		// 預期外的購物車產品
		if val, ok := products[charts[i].ProductId]; !ok {
			result[i].Name = constant.Unknown_Product
			result[i].Amount = 0
			result[i].Image = ""
			result[i].Inventory = 0
			result[i].Status = constant.ProductStatusEnum_Closed
		} else {
			result[i].Name = val.Name
			result[i].Amount = val.Amount
			result[i].Image = val.Image
			result[i].Inventory = val.Inventory
			result[i].Status = val.Status
		}
	}

	return result, nil
}

func (uc *memberChartUseCase) UpdateMemberChart(ctx context.Context, cond *bo.MemberChartUpdateCond) error {
	defer timelogger.LogTime(ctx)()

	memberInfo, ok := ctxs.GetSession(ctx)
	if !ok {
		return errs.MemberTokenError
	}

	if err := uc.validateUpdate(ctx, cond); err != nil {
		return xerrors.Errorf("memberChartUseCase.UpdateMemberChart -> validateUpdate : %w", err)
	}

	db := uc.in.DB.Session()
	if err := uc.in.MemberChartRepo.Update(ctx, db, &po.MemberChartUpdate{
		Id:       cond.Id,
		MemberId: memberInfo.Id,
		Quantity: cond.Quantity,
	}); err != nil {
		return xerrors.Errorf("memberChartUseCase.UpdateMemberChart -> ProductCommon.GetProduct : %w", err)
	}

	return nil
}

func (uc *memberChartUseCase) validateUpdate(ctx context.Context, cond *bo.MemberChartUpdateCond) interface{} {
	if len(cond.Id) == 0 {
		return xerrors.Errorf("memberChartUseCase.UpdateMemberChart -> cond.Id == 0: %w", errs.RequestParamInvalid)
	}

	if cond.Quantity <= 0 {
		return xerrors.Errorf("memberChartUseCase.UpdateMemberChart -> cond.Quantity <= 0: %w", errs.RequestParamInvalid)
	}

	return nil
}

func (uc *memberChartUseCase) CreateMemberChart(ctx context.Context, cond *bo.MemberChartCreateCond) error {
	defer timelogger.LogTime(ctx)()

	memberInfo, ok := ctxs.GetSession(ctx)
	if !ok {
		return errs.MemberTokenError
	}

	if err := uc.validateCreate(ctx, cond); err != nil {
		return xerrors.Errorf("memberChartUseCase.CreateMemberChart -> validateCreate : %w", err)
	}

	db := uc.in.DB.Session()
	// 確認是否已經有此產品
	chart, err := uc.in.MemberChartRepo.First(ctx, db, &po.MemberChartFirst{
		MemberId:  memberInfo.Id,
		ProductId: cond.ProductId,
	})
	if err != nil && errs.ErrDB.NoRow != errs.ConciseParseParse(err) {
		return xerrors.Errorf("memberChartUseCase.CreateMemberChart -> MemberChartRepo.CheckExist : %w", err)
	}

	// chart = nil 代表沒資料 => create   than  update
	if chart == nil {
		if err := uc.in.MemberChartRepo.Insert(ctx, db, &po.MemberChart{
			Id:        uc.in.Uuid.GetUUID(),
			MemberId:  memberInfo.Id,
			ProductId: cond.ProductId,
			Quantity:  cond.Quantity,
		}); err != nil {
			return xerrors.Errorf("memberChartUseCase.CreateMemberChart -> MemberChartRepo.Insert : %w", err)
		}
	} else {
		if err := uc.in.MemberChartRepo.Update(ctx, db, &po.MemberChartUpdate{
			Id:       chart.Id,
			MemberId: memberInfo.Id,
			Quantity: chart.Quantity + cond.Quantity,
		}); err != nil {
			return xerrors.Errorf("memberChartUseCase.CreateMemberChart -> MemberChartRepo.Insert : %w", err)
		}
	}

	return nil
}

func (uc *memberChartUseCase) validateCreate(ctx context.Context, cond *bo.MemberChartCreateCond) interface{} {
	if cond.Quantity <= 0 {
		return xerrors.Errorf("memberChartUseCase.CreateMemberChart -> cond.Quantity <= 0: %w", errs.RequestParamInvalid)
	}

	return nil
}

func (uc *memberChartUseCase) DeleteMemberChart(ctx context.Context, cond *bo.MemberChartDelCond) error {
	defer timelogger.LogTime(ctx)()

	memberInfo, ok := ctxs.GetSession(ctx)
	if !ok {
		return errs.MemberTokenError
	}

	db := uc.in.DB.Session()
	if err := uc.in.MemberChartRepo.Delete(ctx, db, &po.MemberChartDel{
		Id:       cond.Id,
		MemberId: memberInfo.Id,
	}); err != nil {
		return xerrors.Errorf("memberChartUseCase.DeleteMemberChart -> MemberChartRepo.Delete : %w", err)
	}

	return nil
}
