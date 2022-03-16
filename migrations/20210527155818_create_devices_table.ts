import * as Knex from "knex"

export async function up(knex: Knex): Promise<void> {
  return knex.schema.createTable("devices", function (table) {
    table.string("id", 255).primary()
    table.string("name", 255).notNullable()
    table.string("os", 255).notNullable()
    table.string("os_version", 255).notNullable()
    table.string("model", 255).notNullable()
    table.string("uuid", 255).notNullable()
    table.text("token")

    table
      .string("user_id", 255)
      .notNullable()
      .references("id")
      .inTable("users")
      .onDelete("CASCADE")
    table.string("did_address", 255).notNullable()
    table.dateTime("created_at").notNullable()
    table.dateTime("updated_at").notNullable()
    table.dateTime("deleted_at")
  })
}

export async function down(knex: Knex): Promise<void> {
  return knex.schema.dropTableIfExists("devices")
}
