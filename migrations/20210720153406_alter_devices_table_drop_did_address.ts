import * as Knex from "knex"

export async function up(knex: Knex): Promise<void> {
  return knex.schema.alterTable("devices", function (table) {
    table.dropColumn("did_address")
  })
}

export async function down(knex: Knex): Promise<void> {
  return knex.schema.table("devices", function (table) {
    table.string("did_address", 255).notNullable()
  })
}
