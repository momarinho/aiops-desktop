export function createMetricsStream(onUpdate: (data: any) => void, onError?: (error: Error) => void) {
	const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';
	const eventSource = new EventSource(`${API_BASE_URL}/metrics/stream`);

	eventSource.onmessage = (event) => {
		try {
			const data = JSON.parse(event.data);
			onUpdate(data);
		} catch (error) {
			console.error('Failed to parse SSE data:', error);
			onError?.(error as Error);
		}
	};

	eventSource.onerror = (error) => {
		console.error('SSE connection error:', error);
		eventSource.close();
		onError?.(new Error('SSE connection failed'));

		// Auto-reconnect após 5 segundos
		setTimeout(() => {
			console.log('Attempting to reconnect SSE...');
			createMetricsStream(onUpdate, onError);
		}, 5000);
	};

	return eventSource;
}
