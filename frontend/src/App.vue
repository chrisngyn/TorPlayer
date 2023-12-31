<script setup lang="ts"></script>

<template>
  <div class="flex flex-col min-h-screen w-full box-border">
    <nav class="flex justify-between items-center px-4 py-1 bg-stone-900 text-slate-100">
      <ul class="flex">
        <li class="mr-6 p-2">
          <router-link class="hover:text-red-700" to="/">Home</router-link>
        </li>
        <li class="mr-6 p-2">
          <router-link class="hover:text-red-700" to="/about">About</router-link>
        </li>
        <!--        <li class="mr-6 p-2">-->
        <!--          <router-link class="hover:text-red-700" to="/video/123">Video</router-link>-->
        <!--        </li>-->
      </ul>
      <div>
        <button class="px-4 py-2 bg-red-600 rounded" @click="$router.back()">Back</button>
      </div>
    </nav>
    <div class="flex justify-center items-center min-w-full flex-grow bg-stone-800 text-slate-100">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>
li:has(> a.router-link-active) {
  @apply border-b-2 border-red-700
}
</style>
