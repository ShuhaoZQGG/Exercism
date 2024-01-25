class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    return @line.split(": ")[1].strip
  end

  def log_level
    log_level_with_brackets = @line.split(": ")[0]
    log_level_upper_case = log_level_with_brackets.slice(1, log_level_with_brackets.size - 2)
    log_level_lower_case = log_level_upper_case.downcase
  end

  def reformat
    message + " (" + log_level + ")"
  end
end
