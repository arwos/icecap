<script setup lang="ts">
import ISPList from '@/components/oauth/ISPList.vue'
</script>

<template>
  <Suspense>
    <template #default>
      <aside class="box box-light area shadow">
        <ISPList />
      </aside>
    </template>

    <template #fallback>
      <aside class="box box-light area shadow">
        <div>Loading...</div>
      </aside>
    </template>
  </Suspense>
</template>

<style scoped lang="scss">
aside {
  margin-top: 20vh;
  border-radius: var(--ong-radius-size) !important;
}
</style>