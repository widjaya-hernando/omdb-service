'use strict';

module.exports = {
  async up (queryInterface, Sequelize) {
    return queryInterface.createTable('transactions', {
      id: {
        type: Sequelize.BIGINT(64),
        allowNull: false,
        primaryKey: true,
        autoIncrement: true,
      },
      product_id: {
        type: Sequelize.BIGINT(64),
        allowNull: false,
      },
      user_id: {
        type: Sequelize.BIGINT(64),
        allowNull: false,
      },
      quantity: {
        type: Sequelize.STRING,
        allowNull: false,
      },
      total: {
        type: Sequelize.STRING,
        allowNull: false,
      },
      created_at: {
        type: Sequelize.BIGINT,
        allowNull: false,
      },
    })
  },

  async down (queryInterface, Sequelize) {
    return queryInterface.dropTable('transactions');
  }
};
