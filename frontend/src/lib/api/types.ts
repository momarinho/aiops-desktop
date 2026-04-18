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

export interface ExplainAlertContext {
	hostname?: string;
	service?: string;
	recent_events?: string[];
	recent_actions?: string[];
	additional_notes?: string;
}

export interface ExplainAlertRequest {
	alert_id: string;
	context?: ExplainAlertContext;
}

export interface ExplainAlertResponse {
	summary: string;
	probable_cause: string;
	suggested_actions: string[];
}

export interface ApiErrorResponse {
	error: string;
}

// Action Types
export type ActionType = 'kill_process' | 'restart_container' | 'scale_container';
export type ActionStatus = 'pending' | 'success' | 'failed';

export interface Action {
	id: string;
	type: ActionType;
	target: string;
	parameters: Record<string, any>;
	user: string;
	request_time: string;
	start_time?: string;
	end_time?: string;
	status: ActionStatus;
	output?: string;
	error?: string;
	risky: boolean;
}

export interface ExecuteActionRequest {
	type: ActionType;
	target: string;
	parameters?: Record<string, any>;
	user: string;
}

// Process Types
export interface ProcessInfo {
	pid: number;
	name: string;
	user: string;
	cpu_percent: number;
	memory_mb: number;
	create_time: string;
	command: string;
	is_critical: boolean;
	status: 'critical' | 'system' | 'user';
}
