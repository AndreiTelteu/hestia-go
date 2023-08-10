import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';
import devtools from 'solid-devtools/vite';

export default defineConfig({
  plugins: [
    devtools(),
    solidPlugin(),
  ],
  server: {
    port: 80,
  },
  build: {
    target: 'esnext',
  },
});
