import * as Knex from "knex"

export async function up(knex: Knex): Promise<void> {
  return knex.schema.createTable("configs_did", function (table) {
    table.string("did_address", 255).notNullable().unique()
    table.string("public_key_pem", 255).notNullable().unique()
    table.string("private_key_pem", 255).notNullable().unique()
  })
}

export async function down(knex: Knex): Promise<void> {
  return knex.schema.dropTableIfExists("configs_did")
}
