<script setup lang="ts">

</script>

<template>
  <main class="flex flex-col justify-center items-center w-full h-full">
    <h1 class="text-4xl font-bold mb-2">TorPlayer</h1>
    <h3>Xem nhung bo phim yeu thich cua ban</h3>

    <div class="flex flex-col justify-between items-center mt-4">
      <label for="textInput" class="block text-sm mb-2">Input Torrent source</label>
      <input
        type="text"
        id="textInput"
        name="textInput"
        placeholder="Enter magnet, hash or torrent file link"
        class="p-4 rounded border border-gray-400 w-full bg-stone-700 text-stone-100 hover:border-red-700 hover:bg-stone-800"
      />

      <label for="fileInput" class="block text-sm my-2">Or Upload Torrent File</label>
      <input
        type="file"
        id="fileInput"
        name="fileInput"
        accept=".torrent"
        class="border text-stone-100 border-gray-400 hover:border-red-700 file:mr-5 rounded bg-stone-700 file:border-[0px] file:p-4 file:bg-stone-700 file:rounded-l file:text-stone-100 hover:file:cursor-pointer hover:file:bg-stone-800 hover:file:text-red-700"
      />
      <button class="px-4 py-2 my-3 bg-red-600 rounded" @click="$router.push({name:'video', params: {id: '123'}})">
        Play
      </button>
    </div>
  </main>
</template>

