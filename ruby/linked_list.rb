class Node < Struct.new(:value, :previous, :next)
end

class LinkedList
  attr_accessor :head, :tail

end

