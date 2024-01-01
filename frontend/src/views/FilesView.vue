<script setup lang="ts">

import { useRoute } from "vue-router";
import { onMounted, ref } from "vue";
import { main } from "../../wailsjs/go/models";
import { GetTorrentInfo } from "../../wailsjs/go/main/App";


const route = useRoute();

const loading = ref(true);
const torrentInfo = ref<main.TorrentInfo | null>(null);

onMounted(async () => {
  console.log(route.params.infoHash);
  try {
    const infoHash = typeof route.params.infoHash == "string" ? route.params.infoHash : route.params.infoHash[0];
    torrentInfo.value = await GetTorrentInfo(infoHash);
  } finally {
    loading.value = false;
  }
});

const bytesLengthToSize = (length: number): string => {
  const sizes = ["B", "KB", "MB", "GB", "TB"];
  if (length == 0) return "0 B";
  const i = Math.floor(Math.log(length) / Math.log(1024));
  return (length / Math.pow(1024, i)).toFixed(2) + " " + sizes[i];
}

function isVideoFile(fileName: string): boolean {
  const extension = getFileExtension(fileName);
  const videoExtensions = ["mp4", "mov", "avi", "mkv", "mpg", "wmv"];
  return videoExtensions.includes(extension);
}

function getFileExtension(fileName:string): string {
  const parts = fileName.split('.');
  return parts[parts.length - 1].toLowerCase();
}
</script>

<template>
  <main class="w-full min-h-full p-4">
    <p class="text-sm italic">Info Hash: {{ $route.params.infoHash }}</p>
    <div v-if="loading">Loading...</div>
    <template v-else>
      <template v-if="torrentInfo">
        <h3 class="text-4xl my-4">{{ torrentInfo.name }}</h3>
        <p>Size: {{ bytesLengthToSize(torrentInfo.length) }}</p>
        <table class="table-auto w-full my-4 border-2 border-gray-400">
          <thead>
          <tr class="border-gray-400 border-b-2">
            <th class="px-4 py-2">Name</th>
            <th class="px-4 py-2">Length</th>
            <th class="px-4 py-2">Completed</th>
            <th class="px-4 py-2">Action</th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="file in torrentInfo.files" :key="file.displayPath" class="hover:text-red-600">
            <td class="px-4 py-2">{{ file.displayPath }}</td>
            <td class="px-4 py-2 text-right">{{ bytesLengthToSize(file.length) }}</td>
            <td class="px-4 py-2 text-right">{{ (file.bytesCompleted / file.length).toFixed(2) * 100  }}%</td>
            <td class="px-4 py-2 text-center">
              <button v-if="isVideoFile(file.displayPath)" class="px-4 py-2 bg-red-600 rounded text-slate-100">
                <font-awesome-icon icon="fa-play" />
                Watch
              </button>
            </td>
          </tr>
          </tbody>
        </table>
      </template>
    </template>
  </main>
</template>
