<script setup lang="ts">
import { useFetch } from '@vueuse/core'
import ISPItem from '@/components/oauth/ISPItem.vue'
import type { ISPModel } from '@/components/oauth/types'

const url = '/webapi/isp/list.json'
const { error, data } = await useFetch(url).get().json<ISPModel[]>()
</script>

<template>
  <div class="row">
    <div class="col-12">
      <h4>Authentication using OAuth</h4>
      <hr>
    </div>
  </div>
  <div v-if="error">
    <div class="box box-danger">
      {{error}}
    </div>
  </div>
  <div v-else class="row">
    <div class="mr-1" v-for="item in data" :key="item.code">
      <ISPItem v-bind="item"  />
    </div>
  </div>
</template>


<style scoped lang="scss"></style>