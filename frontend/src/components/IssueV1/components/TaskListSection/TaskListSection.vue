<template>
  <div class="relative">
    <div
      ref="taskBar"
      class="task-list gap-2 px-4 py-2 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 3xl:grid-cols-5 4xl:grid-cols-6 overflow-y-auto"
      :class="{
        'more-bottom': taskBarScrollState.bottom,
        'more-top': taskBarScrollState.top,
      }"
      :style="{
        'max-height': `${MAX_LIST_HEIGHT}px`,
      }"
    >
      <TaskCard v-for="(task, i) in taskList" :key="i" :task="task" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from "vue";
import { useVerticalScrollState } from "@/composables/useScrollState";
import { useIssueContext } from "../../logic";
import TaskCard from "./TaskCard.vue";

// 3 * lines of cards + 2 * top and bottom padding + 2 * horizontal gaps + jitter
const MAX_LIST_HEIGHT = 207;

const { selectedStage } = useIssueContext();
const taskBar = ref<HTMLDivElement>();
const taskBarScrollState = useVerticalScrollState(taskBar, MAX_LIST_HEIGHT);

const taskList = computed(() => selectedStage.value.tasks);
</script>

<style scoped lang="postcss">
.task-list::before {
  @apply absolute top-0 h-4 w-full -ml-2 z-10 pointer-events-none transition-shadow;
  content: "";
  box-shadow: none;
}
.task-list::after {
  @apply absolute bottom-0 h-4 w-full -ml-2 z-10 pointer-events-none transition-shadow;
  content: "";
  box-shadow: none;
}
.task-list.more-top::before {
  box-shadow: inset 0 0.3rem 0.25rem -0.25rem rgb(0 0 0 / 10%);
}
.task-list.more-bottom::after {
  box-shadow: inset 0 -0.3rem 0.25rem -0.25rem rgb(0 0 0 / 10%);
}
</style>
