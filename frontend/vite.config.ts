import VueI18nPlugin from "@intlify/unplugin-vue-i18n/vite";
import yaml from "@rollup/plugin-yaml";
import legacy from "@vitejs/plugin-legacy";
import vue from "@vitejs/plugin-vue";
import vueJsx from "@vitejs/plugin-vue-jsx";
import { CodeInspectorPlugin } from "code-inspector-plugin";
import { fileURLToPath, URL } from "node:url";
import { resolve } from "path";
import IconsResolver from "unplugin-icons/resolver";
import Icons from "unplugin-icons/vite";
import Components from "unplugin-vue-components/vite";
import { defineConfig } from "vite";
import viteCompression from "vite-plugin-compression";
import qiankun from "vite-plugin-qiankun";
import { name } from "./package.json";

// import { writeHeapSnapshot } from 'v8';

// setInterval(() => {
//   writeHeapSnapshot(); // 构建过程中定期生成堆快照
// }, 20000);

const SERVER_PORT = parseInt(process.env.PORT ?? "3008", 10) ?? 3008;
const HTTPS_PORT = 443;
const LOCAL_ENDPOINT = "http://172.30.130.130:8082";

// NOTE: the following lines is to solve https://github.com/gitpod-io/gitpod/issues/6719
// tl;dr : the HMR(hot module replacement) will behave differently when VPN is on, and by manually set its port to 443 should prevent this issue.
const IS_RUNNING_GITPOD =
  process.env["GITPOD_WORKSPACE_ID"] !== null &&
  process.env["GITPOD_WORKSPACE_ID"] !== undefined;

const extractHostPort = (url: string) => {
  const parsed = new URL(url);
  return parsed.host;
};

export default defineConfig({
  plugins: [
    legacy({
      targets: ["> 0.08%, not dead"],
      additionalLegacyPolyfills: ["regenerator-runtime/runtime"],
    }),
    vue(),
    qiankun("app/web-dmmp-application", { useDevMode: true }),
    vueJsx(),
    // https://github.com/intlify/vite-plugin-vue-i18n
    VueI18nPlugin({
      include: [resolve(__dirname, "src/locales/**")],
      strictMessage: false,
    }),
    Components({
      allowOverrides: true,
      // auto import icons
      resolvers: [
        IconsResolver({
          prefix: "",
        }),
      ],
    }),
    Icons({
      compiler: "vue3",
    }),
    yaml(),
    CodeInspectorPlugin({
      bundler: "vite",
    }),
    viteCompression({
      algorithm: "gzip",
      ext: ".gz",
      threshold: 10240,
      deleteOriginFile: false,
      filter: /\.(js|css|json|html|ico|svg|png|jpg|jpeg)(\?.*)?$/i,
      compressionOptions: { level: 9 },
      verbose: true,
      disable: false,
    }),
  ],
  css: {
    devSourcemap: false, // 生产构建关闭 Source Map
    preprocessorOptions: {
      scss: {
        additionalData: `@import "./src/variables";`, // 集中导入全局变量
      },
    },
  },
  build: {
    outDir: "dist",
    sourcemap: false,
    minify: "terser",
    terserOptions: {
      compress: {
        drop_console: true,
        drop_debugger: true,
        pure_funcs: ["console.log"],
      },
    },
    rollupOptions: {
      output: {
        format: "es",
        name: name,
        globals: {
          vue: "Vue",
        },
        inlineDynamicImports: false,
        manualChunks(id) {
          // 将大型依赖单独分块
          if (id.includes("node_modules")) {
            if (id.includes("monaco-editor")) {
              return "monaco-editor";
            }
            if (id.includes("vscode")) {
              return "vscode";
            }
            if (id.includes("@iconify/json")) {
              return "iconify";
            }
            if (id.includes("vue") || id.includes("pinia")) {
              return "vue-vendor";
            }
            return "vendors";
          }
          // 组件按目录分块
          if (id.includes("/src/components/")) {
            const match = id.match(/components\/([^/]+)/);
            if (match) {
              return `component-${match[1].toLowerCase()}`;
            }
            return "components";
          }
        },
        entryFileNames: "js/[name]-[hash].js",
        assetFileNames: "[ext]/[name]-[hash].[ext]",
        chunkFileNames: "js/[name]-[hash].js",
      },
    },
  },
  server: {
    port: SERVER_PORT,
    host: "0.0.0.0",
    proxy: {
      "/v1:adminExecute": {
        target: `ws://${extractHostPort(LOCAL_ENDPOINT)}/`,
        changeOrigin: true,
        ws: true,
      },
      "/lsp": {
        target: `ws://${extractHostPort(LOCAL_ENDPOINT)}/`,
        changeOrigin: true,
        ws: true,
      },
      "/api": {
        target: `${LOCAL_ENDPOINT}/api`,
        changeOrigin: true,
        rewrite: (path: string) => path.replace(/^\/api/, ""),
      },
      "/hook": {
        target: LOCAL_ENDPOINT,
        changeOrigin: true,
      },
      "/v1": {
        target: `${LOCAL_ENDPOINT}/v1`,
        changeOrigin: true,
        rewrite: (path: string) => path.replace(/^\/v1/, ""),
      },
      "/bytebase/api": {
        target: `${LOCAL_ENDPOINT}`,
        changeOrigin: true,
        rewrite: (path: string) => path.replace(/^\/bytebase\/api/, ""),
        bypass(req: any, res: any, options: any) {
          if (!req.url) return;
          const proxyURL = options.target + options.rewrite(req.url);
          console.log("x-req-proxyURL", proxyURL);
        },
      },
    },
    hmr: {
      port: IS_RUNNING_GITPOD ? HTTPS_PORT : SERVER_PORT,
    },
  },
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
      "@sql-lsp": fileURLToPath(
        new URL("./src/plugins/sql-lsp", import.meta.url)
      ),
      "@public": fileURLToPath(new URL("./public", import.meta.url)),
    },
    dedupe: ["vscode"],
  },
  envPrefix: "BB_",
  define: {
    _global: {},
  },
  optimizeDeps: {
    include: [
      "vue",
      "vue-router",
      "pinia",
      "axios",
      "monaco-editor/esm/vs/editor/editor.worker.js",
      "monaco-editor/esm/vs/language/json/json.worker.js",
      "monaco-editor/esm/vs/language/css/css.worker.js",
      "monaco-editor/esm/vs/language/html/html.worker.js",
      "monaco-editor/esm/vs/language/typescript/ts.worker.js",
    ],
    exclude: ["@iconify/json"],
  },
});
