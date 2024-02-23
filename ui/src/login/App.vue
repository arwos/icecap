<script setup lang="ts">
import ProvidersForm from '@/login/ProvidersForm.vue'
</script>

<template>
  <Suspense>
    <template #default>
      <aside class="bq bq-success area">
        <ProvidersForm />
      </aside>
    </template>

    <template #fallback>
      <aside class="bq bq-light area">
        <div>Loading...</div>
      </aside>
    </template>
  </Suspense>
</template>

<style scoped lang="scss">
aside {
  margin-top: 20vh;
}
</style>
