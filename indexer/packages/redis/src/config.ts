/**
 * Environment variables required by redis module.
 */

import {
  parseInteger,
  parseString,
  parseSchema,
} from '@nemo-network-indexer/base/build';

export const redisConfigSchema = {
  // Required environment variables.
  REDIS_URL: parseString({
    default: 'redis://redis:6379',
  }),
  REDIS_RECONNECT_TIMEOUT_MS: parseInteger({ default: 500 }),
  REDIS_RECONNECT_ATTEMPT_ERROR_THRESHOLD: parseInteger({ default: 10 }),
};

export default parseSchema(redisConfigSchema);
