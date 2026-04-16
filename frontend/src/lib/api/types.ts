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

export type AlertSeverity = 'warning' | 'critical';
export type AlertStatus = 'firing' | 'acknowledged' | 'silenced' | 'resolved';

export interface Alert {
	id: string;
	severity: AlertSeverity;
	status: AlertStatus;
	description: string;
	metric_name: string;
	threshold: number;
	current_value: number;
	started_at?: string;
	updated_at: string;
	acknowledged_at?: string;
	silenced_at?: string;
	resolved_at?: string;
}
