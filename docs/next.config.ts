import { NextConfig } from 'next';

import path from 'path';

const nextConfig: NextConfig = {
  reactStrictMode: false,
  output: 'export',
  distDir: './out',
  trailingSlash: true,
  images: {
    unoptimized: true,
  },
  generateBuildId: async () => {
    return 'monetr.app';
  },
  transpilePackages: ['@monetr/interface'],
  sassOptions: {
    implementation: 'sass-embedded',
  },
  typescript: {
    // Fuck you next
    ignoreBuildErrors: true,
  },
  webpack: (
    config,
    _,
  ) => {
    // Important: return the modified config
    config.resolve = {
      ...config?.resolve,
      alias: {
        ...config?.resolve?.alias,
        '@monetr/docs': path.resolve(__dirname, 'src'),
        '@monetr/interface': path.resolve(__dirname, '../interface/src'),
      },
      modules: [
        ...config?.resolve?.modules,
        path.resolve(__dirname, '../interface/src'),
      ],
      extensions: [
        ...config?.resolve?.extensions,
        '.svg',
      ],
      extensionAlias: {
        ...config?.resolve?.extensionAlias,
        '.js': ['.ts', '.tsx', '.js', '.jsx'],
        '.mjs': ['.mts', '.mjs'],
        '.cjs': ['.cts', '.cjs'],
        '.svg': ['.svg'],
      },
    };
    return config;
  },
};

const withNextra = require('nextra')({
  theme: 'nextra-theme-docs',
  themeConfig: './theme.config.tsx',
  flexsearch: {
    codeblocks: false,
  },
});

module.exports = withNextra(nextConfig);
