<script setup lang="ts">
import Table from "@/components/Table.vue";
</script>

<template>
  <div class="flex flex-col h-screen bg-gray-100">

    <div class="bg-white text-white shadow w-full p-2 flex items-center justify-between">
      <div class="flex items-center">
        <div class="flex items-center">
          <h2 class="font-bold text-xl text-black">Stock Analysis</h2>
        </div>
      </div>
    </div>

    <div class="flex-1 flex flex-wrap">
      <div class="flex-1 p-4 w-full md:w-1/2">
        <div class="mt-8 bg-white p-4 shadow rounded-lg">
          <h2 class="text-gray-500 text-lg font-semibold pb-4">Stocks</h2>
          <div class="my-1"></div>
         <Table/>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
