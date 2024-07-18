import { createChannel, createClient as createC } from "nice-grpc";
import { UsersClient, UsersDefinition } from "./users";

export class Client {
  users: UsersClient;

  constructor(address = "localhost:3000") {
    const channel = createChannel(address);
    this.users = createC(UsersDefinition, channel);
  }
}
