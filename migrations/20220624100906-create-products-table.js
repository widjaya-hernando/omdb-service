'use strict';

module.exports = {
  async up (queryInterface, Sequelize) {
    return queryInterface.createTable('products', {
      id: {
        type: Sequelize.BIGINT(64),
        allowNull: false,
        primaryKey: true,
        autoIncrement: true,
      },
      product_name: {
        type: Sequelize.STRING,
        allowNull: false,
      },
      quantity: {
        type: Sequelize.BIGINT(64),
        allowNull: false,
      },
      price: {
        type: Sequelize.BIGINT(64),
        allowNull: false,
      },
      created_at: {
        type: Sequelize.BIGINT,
        allowNull: false,
      },
      updated_at: {
        type: Sequelize.BIGINT,
        allowNull: true,
      },
    })
  },

  async down (queryInterface, Sequelize) {
    return queryInterface.dropTable('products');
  }
};
