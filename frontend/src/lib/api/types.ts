export interface Metric {
	name: string;
	value: number;
	unit: string;
	labels?: Record<string, string>;
}

export interface MetricsResponse {
	timestamp: string;
	metrics: Metric[];
}

export interface HealthResponse {
	status: string;
}
