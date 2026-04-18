import { API_BASE_URL } from './config';
import type {
	Action,
	Alert,
	ApiErrorResponse,
	ExecuteActionRequest,
	ExplainAlertRequest,
	ExplainAlertResponse,
	HealthResponse,
	MetricsResponse,
	ProcessInfo
} from './types';
export { createMetricsStream } from './stream';

export async function getHealth(): Promise<HealthResponse> {
	const url = `${API_BASE_URL}/health`;
	console.log('[API] Fetching health from:', url);
	const response = await fetch(url);
	console.log('[API] Health response status:', response.status);
	if (!response.ok) {
		throw new Error(`Health check failed: ${response.statusText}`);
	}
	const data = await response.json();
	console.log('[API] Health data:', data);
	return data;
}

export async function getMetrics(): Promise<MetricsResponse> {
	const url = `${API_BASE_URL}/metrics`;
	console.log('[API] Fetching metrics from:', url);
	const response = await fetch(url);
	console.log('[API] Metrics response status:', response.status);
	if (!response.ok) {
		throw new Error(`Failed to fetch metrics: ${response.statusText}`);
	}
	const data = await response.json();
	console.log('[API] Metrics data:', data);
	return data;
}

export async function getAlerts(): Promise<Alert[]> {
	const url = `${API_BASE_URL}/alerts`;
	console.log('[API] Fetching alerts from:', url);
	const response = await fetch(url);
	console.log('[API] Alerts response status:', response.status);
	if (!response.ok) {
		throw new Error(`Failed to fetch alerts: ${response.statusText}`);
	}
	const data = (await response.json()) as Alert[];
	console.log('[API] Alerts data:', data);
	return data;
}

export async function getAlertById(id: string): Promise<Alert> {
	const url = `${API_BASE_URL}/alerts/${id}`;
	console.log('[API] Fetching alert from:', url);
	const response = await fetch(url);
	console.log('[API] Alert response status:', response.status);
	if (!response.ok) {
		throw new Error(await getApiErrorMessage(response, `Failed to fetch alert: ${response.statusText}`));
	}
	const data = (await response.json()) as Alert;
	console.log('[API] Alert data:', data);
	return data;
}

export async function acknowledgeAlert(id: string): Promise<Alert> {
	return postAlertAction(id, 'acknowledge');
}

export async function silenceAlert(id: string): Promise<Alert> {
	return postAlertAction(id, 'silence');
}

async function postAlertAction(id: string, action: 'acknowledge' | 'silence'): Promise<Alert> {
	const url = `${API_BASE_URL}/alerts/${id}/${action}`;
	console.log(`[API] Posting alert action ${action} to:`, url);
	const response = await fetch(url, { method: 'POST' });
	console.log('[API] Alert action response status:', response.status);
	if (!response.ok) {
		throw new Error(`Failed to ${action} alert: ${response.statusText}`);
	}
	const data = (await response.json()) as Alert;
	console.log('[API] Alert action data:', data);
	return data;
}

export async function explainAlert(request: ExplainAlertRequest): Promise<ExplainAlertResponse> {
	const url = `${API_BASE_URL}/ai/explain-alert`;
	console.log('[API] Requesting alert explanation from:', url, request);
	const response = await fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(request)
	});
	console.log('[API] Explain alert response status:', response.status);
	if (!response.ok) {
		throw new Error(
			await getApiErrorMessage(response, `Failed to explain alert: ${response.statusText}`)
		);
	}
	const data = (await response.json()) as ExplainAlertResponse;
	console.log('[API] Explain alert data:', data);
	return data;
}

// Action API functions
export async function getActions(): Promise<Action[]> {
	const url = `${API_BASE_URL}/actions`;
	console.log('[API] Fetching actions from:', url);
	const response = await fetch(url);
	console.log('[API] Actions response status:', response.status);
	if (!response.ok) {
		throw new Error(`Failed to fetch actions: ${response.statusText}`);
	}
	const data = (await response.json()) as Action[];
	console.log('[API] Actions data:', data);
	return data;
}

async function getApiErrorMessage(response: Response, fallback: string): Promise<string> {
	try {
		const error = (await response.json()) as ApiErrorResponse;
		if (typeof error.error === 'string' && error.error.trim() !== '') {
			return error.error;
		}
	} catch {
		// Ignore parse failures and fall back to the caller-provided message.
	}

	return fallback;
}

export async function getActionById(id: string): Promise<Action> {
	const url = `${API_BASE_URL}/actions/${id}`;
	console.log('[API] Fetching action from:', url);
	const response = await fetch(url);
	console.log('[API] Action response status:', response.status);
	if (!response.ok) {
		throw new Error(`Failed to fetch action: ${response.statusText}`);
	}
	const data = (await response.json()) as Action;
	console.log('[API] Action data:', data);
	return data;
}

export async function executeAction(request: ExecuteActionRequest): Promise<Action> {
	const url = `${API_BASE_URL}/actions`;
	console.log('[API] Executing action:', request);
	const response = await fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(request)
	});
	console.log('[API] Execute action response status:', response.status);
	if (!response.ok) {
		const errorText = await response.text();
		throw new Error(`Failed to execute action: ${response.statusText} - ${errorText}`);
	}
	const data = (await response.json()) as Action;
	console.log('[API] Execute action data:', data);
	return data;
}

// Process API functions
export async function getProcesses(): Promise<ProcessInfo[]> {
	const url = `${API_BASE_URL}/processes`;
	console.log('[API] Fetching processes from:', url);
	const response = await fetch(url);
	console.log('[API] Processes response status:', response.status);
	if (!response.ok) {
		throw new Error(`Failed to fetch processes: ${response.statusText}`);
	}
	const data = (await response.json()) as ProcessInfo[];
	console.log('[API] Processes data:', data);
	return data;
}

export async function getProcessByPID(pid: number): Promise<ProcessInfo> {
	const url = `${API_BASE_URL}/processes/${pid}`;
	console.log('[API] Fetching process from:', url);
	const response = await fetch(url);
	console.log('[API] Process response status:', response.status);
	if (!response.ok) {
		throw new Error(`Failed to fetch process: ${response.statusText}`);
	}
	const data = (await response.json()) as ProcessInfo;
	console.log('[API] Process data:', data);
	return data;
}
