import { API_BASE_URL } from './config';
import type { Alert, HealthResponse, MetricsResponse } from './types';
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
