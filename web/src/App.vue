<script lang="ts">
import { defineComponent } from "vue";
import ListPage from "./components/ListPage.vue";

export default defineComponent({
  name: "MuchTodo",
  components: {
    ListPage,
  },
  data() {
    return {
      tasklists: [],
    };
  },
});
</script>

<template>
  <div id="Title">
    <title>MuchTodo: List and Note Management</title>
    <h1 id="master-title">MuchTodo: List and Note Management</h1>
  </div>
  <div id="MuchTodoRoot">
    <h2>Available Tasks:</h2>
    <ListPage />
  </div>
</template>

<style>
#master-title {
  color: rgba(0, 68, 255, 0.795);
  /* align-content: center; */
  text-align: center;
}
#muchtodo-root {
  text-align: center;
  align-content: center;
  align-items: center;
}
</style>
