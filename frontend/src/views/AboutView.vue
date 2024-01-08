<script setup lang="ts">

import { onMounted, ref } from "vue";
import { GetCurrentConfig } from "../../wailsjs/go/main/App";
import type { torrent } from "../../wailsjs/go/models";

const currentConfig = ref<torrent.Config>();

onMounted(async () => {
  currentConfig.value = await GetCurrentConfig();
});

</script>

<template>
  <main class="w-full p-4">
    <p>Torrent files will be saved at:</p>
    <p class="font-bold">{{ currentConfig?.dataDir }}</p>
    <p>And will be deleted after app closed</p>
  </main>
</template>
