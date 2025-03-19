<script setup lang="ts">
  import {onMounted, ref, watch} from 'vue'

  let headers = [
    "Ticker",
    "Company",
    "Brokerage",
    "Action",
    "Rating From",
    "Rating To",
    "Target From",
    "Target To"
  ]

  const items_response = ref({
    page: 1,
    offset: 1,
    totalPages: 1,
    totalRows: 0,
    rows: []
  })

  let options = [
    { value: "stocks", "label": "All" },
    { value: "recommends-stocks", "label": "Recommends" },
  ]
  let selectedOption = ref(options[0].value)

   let limit = ref(10)
   let page = ref(1)

  let chargeData = async () => {
    fetch(`/api/${selectedOption.value}?limit=${limit.value}&page=${page.value}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    }).then((res) => {
       res.json().then((json) => {
         items_response.value = json
      })
    }).catch(err => {
      console.error(err)
    })
  }

  const goToPage = (newPage: any) => {
    if (newPage >= 1 && newPage <= items_response.value.totalPages) {
      page.value = newPage
    }
  }

  const getPaginationRange = () => {
    const total = items_response.value.totalPages
    const current = page.value
    const delta = 3
    const range = []

    for (let i = Math.max(2, current - delta); i <= Math.min(total - 1, current + delta); i++) {
      range.push(i)
    }

    if (current - delta > 2) {
      range.unshift('...')
    }
    if (current + delta < total - 1) {
      range.push('...')
    }

    range.unshift(1)
    if (total > 1) {
      range.push(total)
    }

    return range
  }

  watch([selectedOption, page, limit], () => {
    chargeData()
  })

  onMounted(
    () => {
      chargeData()
    }
  )
</script>

<template>
  <div class="flex items-center">
    <select v-model="selectedOption" class="ml-2 p-2 border border-gray-300 rounded-lg">
      <option v-for="option in options" :value="option.value">{{ option.label }}</option>
    </select>
  </div>

  <div class="relative flex flex-col w-full h-full overflow-scroll text-gray-700 bg-white shadow-md rounded-lg bg-clip-border">
    <table class="w-full table-auto text-sm  min-w-max">
      <thead>
        <tr class="text-sm leading-normal">
          <th class="py-2 px-4 bg-grey-lightest font-bold uppercase text-sm text-grey-light border-b border-grey-light" v-for="header in headers">
            {{header}}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr class="hover:bg-grey-lighter" v-if="items_response.rows?.length == 0">
          <td :colspan="headers.length" class="py-2 px-4 border-b border-grey-light text-center">Cargando...</td>
        </tr>
        <tr class="hover:bg-grey-lighter" v-if="items_response.rows?.length > 0" v-for="item in items_response.rows">
          <td class="py-2 px-4 border-b border-grey-light text-center" v-for="header in headers">
            {{
              ["TargetFrom", "TargetTo"].includes(
                header.replace(" ", "")
              )
                ? `$${item[header.replace(" ", "")]}`
                :
              item[header.replace(" ", "")]
            }}
          </td>
        </tr>
      </tbody>
    </table>

    <div class="flex justify-between items-center px-4 py-3">
      <div class="text-sm text-slate-500">
        Showing <b>{{items_response.offset + 1}}-{{items_response.offset + limit}}</b> of {{items_response.totalRows}}
      </div>
      <div class="flex space-x-1">
        <button :disabled="items_response.page <= 1" @click="goToPage(page - 1)" class="px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 bg-white border border-slate-200 rounded hover:bg-slate-50 hover:border-slate-400 transition duration-200 ease">
          Prev
        </button>
        <button :disabled="number == page || number == '...'" v-for="number in getPaginationRange()" :key="number" @click="goToPage(number)" :class="{'bg-slate-800 text-white': number === page, 'bg-white text-slate-500 hover:bg-slate-50 hover:border-slate-400': number !== page}" class="px-3 py-1 min-w-9 min-h-9 text-sm font-normal border rounded transition duration-200 ease">
          {{ number }}
        </button>
        <button @click="goToPage(page + 1)" class="px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 bg-white border border-slate-200 rounded hover:bg-slate-50 hover:border-slate-400 transition duration-200 ease">
          Next
        </button>
      </div>
    </div>
  </div>

</template>

<style scoped>
</style>
