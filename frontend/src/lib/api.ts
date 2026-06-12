import { Todo } from '@/types/todo';

const BASE = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8989';

export async function listTodos(): Promise<Todo[]> {
  const res = await fetch(`${BASE}/todos`);
  if (!res.ok) throw new Error('list failed');
  return res.json();
}

export async function createTodos(content: string): Promise<Todo> {
  const res = await fetch(`${BASE}/todos`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ content }),
  });
  if (!res.ok) throw new Error('create failed');

  return res.json();
}

export async function toggleTodo(id: number, done: boolean): Promise<Todo> {
  const res = await fetch(`${BASE}/todos/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ done }),
  });
  if (!res.ok) throw new Error('update failed');
  return res.json();
}
