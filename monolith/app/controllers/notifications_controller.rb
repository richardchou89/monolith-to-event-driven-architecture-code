class NotificationsController < ApplicationController
  def test
    SnsPublishJob.perform_later(
      message: {
        options: options,
        clientId: @client_id,
        apiKey: @api_key
      },
      message_attributes: { eventType: { data_type: 'String', string_value: 'send_email' } }
    )
  end
end
