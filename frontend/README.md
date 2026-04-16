# AI-Ops Desktop - Frontend

Frontend application for AI-Ops Desktop, built with Svelte and TypeScript.

## Technology Stack
- **Framework:** SvelteKit
- **Language:** TypeScript
- **Styling:** TailwindCSS with custom theme
- **Icons:** Material Symbols Outlined
- **API Communication:** Fetch + Server-Sent Events (SSE)

## Development

```bash
# Install dependencies
npm install

# Start development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

## Environment Variables

Create a `.env` file in the root directory:

```bash
VITE_API_BASE_URL=http://localhost:8080
```

## Project Structure

```
src/
├── lib/
│   ├── api/
│   │   ├── config.ts       # API configuration
│   │   ├── index.ts        # API client functions
│   │   ├── stream.ts       # SSE client implementation
│   │   └── types.ts        # TypeScript interfaces
│   ├── components/
│   │   └── AppShell.svelte # Main app layout
│   └── index.ts            # Library exports
├── routes/
│   ├── +layout.svelte      # Root layout
│   ├── +page.svelte        # Dashboard page
│   ├── alerts/             # Alerts pages
│   ├── assistant/          # AI assistant pages
│   ├── history/            # History pages
│   └── settings/           # Settings pages
└── routes/layout.css       # Global styles
```

## API Integration

### REST API
```typescript
import { getHealth, getMetrics } from '$lib/api';

const health = await getHealth();
const metrics = await getMetrics();
```

### Server-Sent Events (SSE)
```typescript
import { createMetricsStream } from '$lib/api';

const eventSource = createMetricsStream(
  (data) => {
    // Handle real-time updates
    console.log('New metrics:', data);
  },
  (error) => {
    // Handle connection errors
    console.error('SSE error:', error);
  }
);
```

## Design System

### Color Theme
- **Primary:** Teal (#6feee1)
- **Secondary:** Gray (#c3c6d1)
- **Tertiary:** Orange (#ffd1af)
- **Error:** Red (#ffb4ab)
- **Background:** Dark (#111319)

### Components
- **Panels:** `.panel`, `.panel-soft`, `.panel-deep`
- **Cards:** `.metric-card` with accent colors
- **Status:** `.status-dot` for connection status
- **Grid:** `.grid-overlay` for background patterns

## Development Notes

- The app connects to a Go backend running on port 8080 by default
- Real-time metrics are delivered via SSE at `/metrics/stream`
- The interface uses a minimalist design without glow effects
- All API calls are centralized in the `lib/api` directory
- Type safety is maintained across the entire application

## Building for Production

The frontend is designed to be packaged with Electron for desktop distribution. When building for production:

1. Ensure the backend API URL is correctly configured
2. Run `npm run build` to create optimized assets
3. The Electron wrapper will handle the backend process and serve these files

## Troubleshooting

**CORS Errors:** Ensure the Go backend has CORS enabled for development

**SSE Connection Issues:** Check that `/metrics/stream` endpoint is accessible and the backend is running

**Build Errors:** Clear the `.svelte-kit` folder and try rebuilding

---

**Last Updated:** April 16, 2026
**Status:** Sprint 2 Complete - Live Metrics Pipeline
