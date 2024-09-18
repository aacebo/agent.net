export interface Agent {
  readonly id: string;
  readonly parent_id?: string;
  readonly container_id?: string;
  readonly name: string;
  description?: string;
  status: 'up' | 'down';
  instructions?: string;
  address?: string;
  settings: {
    api_key: string;
    model: string;
    frequency_penalty?: number;
    logit_bias?: Record<string, any>;
    logprobs?: boolean;
  };
  readonly created_at: Date | string;
  readonly updated_at: Date | string;
}
