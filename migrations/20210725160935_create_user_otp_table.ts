import * as Knex from "knex"

export async function up(knex: Knex): Promise<void> {
  return knex.schema.createTable("user_otps", function (table) {
    table.string("id", 255).primary()
    table.string("otp_number", 255).notNullable()
    table
      .string("user_id", 255)
      .notNullable()
      .references("id")
      .inTable("users")
      .onDelete("CASCADE")
    table.dateTime("verified_at")
    table.dateTime("expired_at")
    table.dateTime("revoked_at")
    table.dateTime("created_at").notNullable()
    table.dateTime("updated_at").notNullable()
  })
}

export async function down(knex: Knex): Promise<void> {
  return knex.schema.dropTableIfExists("user_otps")
}
