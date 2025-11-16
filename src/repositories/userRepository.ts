import { User, CreateUserDTO, UpdateUserDTO, Status } from "../models/user";
import { randomUUID } from "crypto";

export class UserRepository {
  private users: Map<string, User> = new Map();

  async create(payload: CreateUserDTO): Promise<User> {
    // check unique email
    for (const u of this.users.values()) {
      if (u.email.toLowerCase() === payload.email.toLowerCase()) {
        throw new Error("email_already_registered");
      }
    }

    const id = randomUUID();
    const now = new Date().toISOString();

    const user: User = {
      id,
      email: payload.email,
      name: payload.name,
      info: payload.info ?? null,
      status: payload.status ?? Status.ACTIVE,
      createdAt: now,
      updatedAt: null,
    };

    this.users.set(id, user);
    return user;
  }

  async list(): Promise<User[]> {
    return Array.from(this.users.values());
  }

  async getById(id: string): Promise<User | null> {
    return this.users.get(id) ?? null;
  }

  async update(id: string, payload: UpdateUserDTO): Promise<User | null> {
    const existing = this.users.get(id);
    if (!existing) return null;

    // If updating email - check uniqueness
    if (payload.email && payload.email.toLowerCase() !== existing.email.toLowerCase()) {
      for (const u of this.users.values()) {
        if (u.email.toLowerCase() === payload.email.toLowerCase()) {
          throw new Error("email_already_registered");
        }
      }
    }

    const now = new Date().toISOString();

    const updated: User = {
      ...existing,
      email: payload.email ?? existing.email,
      name: payload.name ?? existing.name,
      info: payload.info ?? existing.info,
      status: payload.status ?? existing.status,
      updatedAt: now,
    };

    this.users.set(id, updated);
    return updated;
  }

  async delete(id: string): Promise<boolean> {
    return this.users.delete(id);
  }

  // For convenience in tests
  async clear() {
    this.users.clear();
  }
}
