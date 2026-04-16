export type MetricType = 'cpu' | 'memory' | 'disk' | 'network';

export interface Metric {
	type: MetricType;
	name: string;
	value: number;
	unit: string;
	timestamp: string;
	labels?: Record<string, string>;
}

export interface Snapshot {
	timestamp: string;
	metrics: Metric[];
}

export interface MetricsResponse extends Snapshot {}

export interface HealthResponse {
	status: string;
}
