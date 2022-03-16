import * as Knex from "knex";


export async function up(knex: Knex): Promise<void> {
    return knex.schema.createTable('users', function (table) {
        table.string('id', 255).primary()
        table.string('id_card_no', 255).notNullable().unique()
        table.string('first_name', 255).notNullable()
        table.string('last_name', 255).notNullable()
        table.string('did_address', 255)
        table.dateTime('created_at').notNullable()
        table.dateTime('updated_at').notNullable()
        table.dateTime('deleted_at')
    })
}

export async function down(knex: Knex): Promise<void> {
    return knex.schema.dropTableIfExists('users')
}
