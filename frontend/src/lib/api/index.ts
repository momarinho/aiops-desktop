import { API_BASE_URL } from './config';
import type { HealthResponse, MetricsResponse } from './types';

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
