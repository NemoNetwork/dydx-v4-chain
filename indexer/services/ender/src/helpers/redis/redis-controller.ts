import { redis as redisLib } from '@nemo-network-indexer/redis/build/redis/src';
import {
  RedisClient,
} from 'redis';

import config from '../../config';

const res: {
  client: RedisClient,
  connect: () => Promise<void>,
} = redisLib.createRedisClient(config.REDIS_URL, config.REDIS_RECONNECT_TIMEOUT_MS);

export const redisClient: RedisClient = res.client;
export const connect = res.connect;
