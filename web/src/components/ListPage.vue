<script lang="js">

let id = 0;

export default {
  data() {
    return {
      taskitems: [
        {id: id++, 'text': 'Task item 1', done: false, deleted: false}
    ],
      num: 1
    }
  }
}
</script>

<template>
  <p v-for="task in taskitems" :key="task.text">{{ task.text }}</p>
</template>

<style>
.temp {
  color: red;
}
</style>
