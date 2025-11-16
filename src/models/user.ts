export enum Status {
  ACTIVE = "active",
  INACTIVE = "inactive",
  SUSPENDED = "suspended",
}

export interface User {
  id: string; // uuid
  email: string;
  name: string;
  info?: Record<string, unknown> | null;
  status: Status;
  createdAt: string; // ISO date
  updatedAt?: string | null; // ISO date
}

export interface CreateUserDTO {
  email: string;
  name: string;
  info?: Record<string, unknown> | null;
  status?: Status;
}

export interface UpdateUserDTO {
  email?: string;
  name?: string;
  info?: Record<string, unknown> | null;
  status?: Status;
}
