<script setup lang="ts">
import { ref } from 'vue';
import type { Node, Edge } from '@vue-flow/core';
import { VueFlow, useVueFlow } from '@vue-flow/core'

const { onPaneReady } = useVueFlow();
onPaneReady((instance) => instance.fitView());

const nodes = ref<Node[]>([
  {
    id: '1',
    type: 'input',
    position: { x: 250, y: 5 },
    data: { label: 'Node 1' },
  },
  {
    id: '2',
    position: { x: 100, y: 100 },
    data: { label: 'Node 2' },
  },
  {
    id: '3',
    type: 'output',
    position: { x: 400, y: 200 },
    data: { label: 'Node 3' },
  },
  {
    id: '4',
    position: { x: 400, y: 200 },
    data: {
      label: 'Node 4',
      hello: 'world',
    },
  },
]);

const edges = ref<Edge[]>([
  {
    id: 'e1->2',
    source: '1',
    target: '2',
  },
  {
    id: 'e2->3',
    source: '2',
    target: '3',
    animated: true,
  },
  {
    id: 'e3->4',
    source: '3',
    target: '4',
    data: {
      hello: 'world',
    }
  },
]);
</script>

<template>
  <div style="height: 100vh; width: 100vw">
    <VueFlow
      :nodes="nodes"
      :edges="edges"
      :fit-view-on-init="true"
    >
    </VueFlow>
  </div>
</template>

<style>
@import '@vue-flow/core/dist/style.css';
@import '@vue-flow/core/dist/theme-default.css';
</style>
