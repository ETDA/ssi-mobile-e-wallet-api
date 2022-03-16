import * as Knex from "knex"

export async function up(knex: Knex): Promise<void> {
  return knex.schema.table("users", function (table) {
    table.string("email", 255).notNullable()
    table.string("did_address", 255).nullable().alter()
  })
}

export async function down(knex: Knex): Promise<void> {
  return knex.schema.table("users", function (table) {
    table.dropColumn("email")
    table.string("did_address", 255).notNullable().alter()
  })
}
