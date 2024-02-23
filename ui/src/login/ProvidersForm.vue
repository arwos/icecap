<script setup lang="ts">
import { useFetch } from '@vueuse/core'
import type { TProvider } from '@/login/types'
import ProviderButton from '@/login/ProviderButton.vue'

const url = '/webapi/login.json'
const { error, data } = await useFetch(url).get().json<TProvider[]>()
</script>

<template>
  <div class="row">
    <div class="col-12">
      <h4>Sign in with:</h4>
    </div>
  </div>
  <div v-if="error">
    <div class="box box-danger">
      {{ error }}
    </div>
  </div>
  <div v-else class="row">
    <div class="c btn-size" v-for="item in data" :key="item.code">
      <ProviderButton v-bind="item" />
    </div>
  </div>
</template>

<style scoped lang="scss">
.btn-size {
  width: 40px !important;
}
</style>
