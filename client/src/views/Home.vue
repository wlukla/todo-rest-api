<template>
  <div class="home">
    <h1>Home page</h1>

    <ul v-if="todos.length > 0" class="todo-container">
      <li v-for="todo in todos" :key="todo.id">{{ todo.title }}</li>
    </ul>
    <p v-else>No todos for today!</p>

    <p>TEST</p>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  name: "Home",
  data() {
    return {
      todos: [
        { id: "0", title: "Todo #1", checked: false },
        { id: "1", title: "Todo #2", checked: true },
        { id: "2", title: "Todo #3", checked: false }
      ]
    };
  }
});
</script>

<style scoped>
.todo-container {
  margin: 0;
  list-style: none;
}
</style>
