<script setup lang="ts">
import { useFetch } from '@vueuse/core'
import type { TField, TSave } from '@/profile/fields/types'
import FieldText from '@/profile/fields/FieldText.vue'
import FieldInput from '@/profile/fields/FieldInput.vue'
import FieldSelect from '@/profile/fields/FieldSelect.vue'
import { onMounted, ref } from 'vue'

const url = '/webapi/profile.json'
const changes = new Map<string, string>()
const hasChange = ref(false)
const _data = ref<TField[]>([])
const _error = ref<any>()

const setChange = (field: string, value: string) => {
  changes.set(field, value)
  hasChange.value = true
}
const save = async () => {
  const body: TSave[] = []
  changes.forEach((value, field) => body.push({ value, field }))
  const { error } = await useFetch(url).post(body, 'json')
  if (error.value !== null) {
    _error.value = error.value
  } else {
    changes.clear()
    hasChange.value = false
  }
}

onMounted(async () => {
  const { error, data } = await useFetch(url).get().json<TField[]>()
  if (error.value !== null) {
    _error.value = error.value
    return
  }
  _data.value = data.value ?? []
})
</script>

<template>
  <div class="row pos-sticky bg-white">
    <div class="col-12">
      <div class="row">
        <div class="col-10">
          <h4>Profile</h4>
        </div>
        <div class="col-2">
          <button v-if="hasChange" @click="save()" class="btn btn-success mt-1">Save</button>
        </div>
      </div>
      <div v-if="_error" class="bq bq-danger">
        {{ _error }}
      </div>
    </div>
  </div>

  <div class="row" v-for="item in _data" :key="item.field">
    <div class="col-12" v-if="item.type === 'text'">
      <FieldText v-bind="item" />
    </div>
    <div class="col-12" v-if="item.type === 'input'">
      <FieldInput v-bind="item" v-on:value="setChange" />
    </div>
    <div class="col-12" v-if="item.type === 'select'">
      <FieldSelect v-bind="item" v-on:value="setChange" />
    </div>
  </div>
</template>

<style scoped lang="scss">
.bg-white {
  background-color: #ffffff;
}
</style>
