module.exports = {
  apps: [
    {
      name: 'api-chatgpt',
      script: './main',
      watch: false,
      max_memory_restart: '50M',
      restart_delay: 5000,
    },
  ],
};
