import { baseConfigSchema, parseSchema } from '@dydxprotocol-indexer/base';
import { kafkaConfigSchema } from '@dydxprotocol-indexer/kafka';
import { postgresConfigSchema } from '@dydxprotocol-indexer/postgres';
import { redisConfigSchema } from '@dydxprotocol-indexer/redis';

export const configSchema = {
  ...baseConfigSchema,
  ...postgresConfigSchema,
  ...kafkaConfigSchema,
  ...redisConfigSchema,
};

export default parseSchema(configSchema);
