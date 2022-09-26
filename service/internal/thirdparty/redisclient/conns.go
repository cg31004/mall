package redisclient

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/xerrors"

	"simon/mall/service/internal/utils/converter"
)

//go:generate mockery --name IRedisClient --structname MockRedisClient --output mock_redis --outpkg mock_redis --filename mock_redis.go --with-expecter


var (
	once sync.Once
	self *RedisClient
)

func initWithConfig(in digIn) IRedisClient {
	self = &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     in.OpsConf.GetRedisServiceConfig().Address,
			Password: in.OpsConf.GetRedisServiceConfig().Password,
			DB:       in.OpsConf.GetRedisServiceConfig().DB,
		}),
	}

	return self
}

type IRedisClient interface {
	CheckIfKeyExists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	SetJSON(ctx context.Context, key string, value interface{}, duration time.Duration) error
	GetString(ctx context.Context, key string) (string, error)
	GetJSON(ctx context.Context, key string, val interface{}) error
	GetJSONWithMultiKeys(ctx context.Context, val interface{}, keys ...string) error
	SubscribeChannel(ctx context.Context, channels ...string) <-chan *redis.Message
	AddToSet(ctx context.Context, key string, values ...interface{}) error
	AddToSortedSet(ctx context.Context, key string, score float64, value interface{}) error
	GetAllInSet(ctx context.Context, key string) ([]string, error)
	GetAllInSortedSet(ctx context.Context, key string) ([]string, error)
	GetAllInSortedSetWithScores(ctx context.Context, key string, start, end int64) ([]string, error)
	RemoveFromSet(ctx context.Context, key string, values ...interface{}) (int64, error)
	RemoveFromSortedSet(ctx context.Context, key string, value interface{}) error
	RemoveAllFromSet(ctx context.Context, key string) (int64, error)
	RemoveAllFromSortedSet(ctx context.Context, key string) (int64, error)
	CheckInSet(ctx context.Context, key string, value interface{}) (bool, error)
	CheckInSortedSet(ctx context.Context, key string, value string) (bool, error)
	GetSetCount(ctx context.Context, key string) (int64, error)
	Increase(ctx context.Context, key string) (int64, error)
	Decrease(ctx context.Context, key string) (int64, error)
	Publish(ctx context.Context, key string, value interface{}) (int64, error)
	HGet(ctx context.Context, key, field string) (string, error)
	HMGet(ctx context.Context, key string, values []string) (map[string]string, error)
	XAdd(ctx context.Context, topic string, values map[string]interface{}) error
	XGroupCreateMkStream(ctx context.Context, topic, group string) error
	XReadGroup(ctx context.Context, group, consumerID, topic string) ([]redis.XStream, error)
	XAck(ctx context.Context, group, topic string, ids ...string) error
	XDel(ctx context.Context, topic string, ids ...string) error
}

type RedisClient struct {
	client *redis.Client
}

func (client *RedisClient) CheckIfKeyExists(ctx context.Context, key string) (bool, error) {
	val, err := client.client.Exists(ctx, key).Result()
	if err != nil {
		return false, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}

	return val > 0, nil
}

func (client *RedisClient) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	if _, err := client.client.Set(ctx, key, value, duration).Result(); err != nil {
		return xerrors.Errorf("無法從 Redis 新增值 %v 進 Set %s: %w", value, key, err)
	}
	return nil
}

func (client *RedisClient) SetJSON(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	jsonVal, err := json.Marshal(value)
	if err != nil {
		return xerrors.Errorf("無法轉換 JSON 資料 %s: %w", key, err)
	}

	if _, err := client.client.Set(ctx, key, jsonVal, duration).Result(); err != nil {
		return xerrors.Errorf("無法從 Redis 新增值 %v 進 Set %s: %w", value, key, err)
	}
	return nil
}

func (client *RedisClient) GetString(ctx context.Context, key string) (string, error) {
	data, err := client.client.Get(ctx, key).Result()
	if err != nil {
		return "", xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}

	return data, nil
}

func (client *RedisClient) GetJSON(ctx context.Context, key string, val interface{}) error {
	data, err := client.client.Get(ctx, key).Bytes()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}

	err = json.Unmarshal(data, val)
	if err != nil {
		return xerrors.Errorf("無法解析 JSON 資料 %s: %w", key, err)
	}

	return nil
}

func (client *RedisClient) GetJSONWithMultiKeys(ctx context.Context, val interface{}, keys ...string) error {
	data, err := client.client.MGet(ctx, keys...).Result()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 一次取得多筆資料 %v: %w", keys, err)
	}

	rawStrings := convertInterfaceArrayToStringArray(data)
	rawString := "[" + strings.Join(rawStrings, ",") + "]"

	if err := json.Unmarshal([]byte(rawString), val); err != nil {
		return xerrors.Errorf("無法解析 JSON 資料 %s: %w", keys, err)
	}

	return nil
}

func (client *RedisClient) SubscribeChannel(ctx context.Context, channels ...string) <-chan *redis.Message {
	ch := client.client.Subscribe(ctx, channels...).ChannelSize(1000)

	go func() {
		for {
			if len(ch) > 0 {
				fmt.Printf("channel %v: %d\n", channels, len(ch))
			}
			time.Sleep(1 * time.Second)
		}
	}()

	return ch
}

func (client *RedisClient) AddToSet(ctx context.Context, key string, values ...interface{}) error {
	_, err := client.client.SAdd(ctx, key, values...).Result()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 新增值 %v 進 Set %s: %w", values, key, err)
	}
	return err
}

func (client *RedisClient) AddToSortedSet(ctx context.Context, key string, score float64, value interface{}) error {
	_, err := client.client.ZAdd(ctx, key, &redis.Z{Score: score, Member: value}).Result()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 新增值 %v 進 SortedSet %s: %w", value, key, err)
	}
	return err
}

func (client *RedisClient) GetAllInSet(ctx context.Context, key string) ([]string, error) {
	values, err := client.client.SMembers(ctx, key).Result()
	if err != nil {
		return nil, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}
	return values, nil
}

func (client *RedisClient) GetAllInSortedSet(ctx context.Context, key string) ([]string, error) {
	values, err := client.client.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}
	return values, nil
}

func (client *RedisClient) GetAllInSortedSetWithScores(ctx context.Context, key string, start, end int64) ([]string, error) {
	values, err := client.client.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Min: converter.ConvertInt64ToStr(start),
		Max: converter.ConvertInt64ToStr(end),
	}).Result()

	if err != nil {
		return nil, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}
	return values, nil
}

func (client *RedisClient) RemoveFromSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	count, err := client.client.SRem(ctx, key, values...).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法從 Redis 的 Set %s 刪除 %v: %w", key, values, err)
	}
	return count, err
}

func (client *RedisClient) RemoveFromSortedSet(ctx context.Context, key string, value interface{}) error {
	_, err := client.client.ZRem(ctx, key, value).Result()
	if err != nil {
		return xerrors.Errorf("無法從 Redis 的 SortedSet %s 刪除 %v: %w", key, value, err)
	}
	return err
}

func (client *RedisClient) RemoveAllFromSet(ctx context.Context, key string) (int64, error) {
	count, err := client.client.Del(ctx, key).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法從 Redis 刪除 Set %s 的全部值: %w", key, err)
	}
	return count, nil
}

func (client *RedisClient) RemoveAllFromSortedSet(ctx context.Context, key string) (int64, error) {
	count, err := client.client.Del(ctx, key).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法從 Redis 刪除 Sorted Set %s 的全部值: %w", key, err)
	}
	return count, nil
}

func (client *RedisClient) CheckInSet(ctx context.Context, key string, value interface{}) (bool, error) {
	existed, err := client.client.SIsMember(ctx, key, value).Result()
	if err != nil {
		return false, xerrors.Errorf("無法從 Redis 確認 %s 是否存在: %w", key, err)
	}
	return existed, nil
}

func (client *RedisClient) CheckInSortedSet(ctx context.Context, key string, value string) (bool, error) {
	idx, err := client.client.ZRank(ctx, key, value).Result()
	if err != nil {
		return false, xerrors.Errorf("無法從 Redis 確認 %s 是否存在: %w", key, err)
	}
	existed := idx != -1
	return existed, nil
}

func (client *RedisClient) GetSetCount(ctx context.Context, key string) (int64, error) {
	count, err := client.client.SCard(ctx, key).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}
	return count, nil
}

func (client *RedisClient) Increase(ctx context.Context, key string) (int64, error) {
	value, err := client.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法在 Redis 的 %s 加上  1", key, err)
	}
	return value, nil
}

func (client *RedisClient) Decrease(ctx context.Context, key string) (int64, error) {
	value, err := client.client.Decr(ctx, key).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法在 Redis 的 %s 減掉  1", key, err)
	}
	return value, nil
}

func (client *RedisClient) Publish(ctx context.Context, key string, value interface{}) (int64, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return 0, xerrors.Errorf("Marshal json 失敗: %w", err)
	}

	row, err := client.client.Publish(ctx, key, data).Result()
	if err != nil {
		return 0, xerrors.Errorf("無法推送Channel %s 資料是 %v: %s", key, value, err)
	}

	return row, nil
}

func (client *RedisClient) HGet(ctx context.Context, key, field string) (string, error) {
	data, err := client.client.HGet(ctx, key, field).Result()
	if err != nil {
		return "", xerrors.Errorf("無法從 Redis 取得 Key %s 值 %w", key, err)
	}

	return data, nil
}

func (client *RedisClient) HMGet(ctx context.Context, key string, fields []string) (map[string]string, error) {
	data, err := client.client.HMGet(ctx, key, fields...).Result()
	if err != nil {
		return nil, xerrors.Errorf("無法從 Redis 取得 %s: %w", key, err)
	}

	outputData := map[string]string{}
	for i, v := range data {
		if s, ok := v.(string); ok {
			outputData[fields[i]] = s
		}
	}

	return outputData, nil
}

func (client *RedisClient) XAdd(ctx context.Context, topic string, values map[string]interface{}) error {
	args := &redis.XAddArgs{
		Stream:       topic,
		MaxLen:       0,
		MaxLenApprox: 0,
		Values:       values,
	}

	cmd := client.client.XAdd(ctx, args)
	if err := cmd.Err(); err != nil {
		return xerrors.Errorf("無法從 Redis 添加 Stream，資料 = %+v，錯誤 = %w", cmd.String(), err)
	}

	return nil
}

func (client *RedisClient) XGroupCreateMkStream(ctx context.Context, topic, group string) error {
	cmd := client.client.XGroupCreateMkStream(ctx, topic, group, "0")
	if err := cmd.Err(); err != nil {
		if err.Error() == "BUSYGROUP Consumer Group name already exists" {
			return nil
		}
		return xerrors.Errorf("無法從 Redis 創建 Group，cmd = %s，錯誤 = %w", cmd.String(), err)
	}

	return nil
}

func (client *RedisClient) XReadGroup(ctx context.Context, group, consumerID, topic string) ([]redis.XStream, error) {
	args := &redis.XReadGroupArgs{
		Group:    group,
		Consumer: consumerID,
		Streams:  []string{topic, ">"},
		Count:    2,
		Block:    0,
		NoAck:    false,
	}

	res, err := client.client.XReadGroup(ctx, args).Result()
	if err != nil {
		return nil, xerrors.Errorf("無法從 Redis 讀取 Stream，stream = %s，錯誤 = %w", topic, err)
	}

	return res, nil
}

func (client *RedisClient) XAck(ctx context.Context, group, topic string, ids ...string) error {
	if _, err := client.client.XAck(ctx, topic, group, ids...).Result(); err != nil {
		return xerrors.Errorf("無法從 Redis 讀取 Stream，stream = %s，錯誤 = %w", topic, err)
	}

	return nil
}

func (client *RedisClient) XDel(ctx context.Context, topic string, ids ...string) error {
	if _, err := client.client.XDel(ctx, topic, ids...).Result(); err != nil {
		return xerrors.Errorf("無法從 Redis 讀取 Stream，stream = %s，錯誤 = %w", topic, err)
	}

	return nil
}
