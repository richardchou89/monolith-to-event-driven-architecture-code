class SnsPublishJob < ApplicationJob
  queue_as :low

  def perform(message:, message_attributes:)
    SnsService.publish(
      message: message,
      message_attributes: message_attributes
    )
  end
end
