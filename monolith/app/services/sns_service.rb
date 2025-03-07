class SnsService
  def self.publish(message: {}, message_attributes: {})
    return unless message.instance_of?(Hash) && message_attributes.instance_of?(Hash)

    sns_client.publish(
      topic_arn: ENV['AWS_SNS_ARN'],
      message: message.to_json,
      message_attributes: message_attributes
    )
  end

  def self.sns_client
    Aws::SNS::Client.new(
      region: ENV['AWS_SNS_REGION'],
      credentials: Aws::Credentials.new(ENV['AWS_SNS_ACCESS_KEY_ID'], ENV['AWS_SNS_SECRET_ACCESS_KEY'])
    )
  end
end
