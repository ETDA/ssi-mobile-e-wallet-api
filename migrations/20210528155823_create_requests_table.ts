import * as Knex from 'knex'

export async function up (knex: Knex): Promise<void> {
  return knex.schema.createTable('requests', function (table) {
    table.string('id', 255).primary()
    table.string('request_id', 255).notNullable().unique()
    table.text('request_data').notNullable()
    table.string('schema_type', 255).notNullable()
    table.string('credential_type', 255).notNullable()
    table.string('signer', 255).notNullable()
    table.string('requester', 255).notNullable()
    table.string('status', 255).notNullable()
    table.dateTime('created_at').notNullable()
    table.dateTime('updated_at').notNullable()
    table.dateTime('deleted_at')
  })
}

export async function down (knex: Knex): Promise<void> {
  return knex.schema.dropTableIfExists('requests')
}
