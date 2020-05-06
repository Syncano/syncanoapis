# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: syncano/codebox/lb/v1/workerplug.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='syncano/codebox/lb/v1/workerplug.proto',
  package='syncano.codebox.lb.v1',
  syntax='proto3',
  serialized_options=b'Z>github.com/Syncano/syncanoapis/gen/go/syncano/codebox/lb/v1;lb',
  serialized_pb=b'\n&syncano/codebox/lb/v1/workerplug.proto\x12\x15syncano.codebox.lb.v1\"v\n\x17\x43ontainerRemovedRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x14\n\x0c\x63ontainer_id\x18\x02 \x01(\t\x12\x13\n\x0bsource_hash\x18\x03 \x01(\t\x12\x13\n\x0b\x65nvironment\x18\x04 \x01(\t\x12\x0f\n\x07user_id\x18\x05 \x01(\t\"\x1a\n\x18\x43ontainerRemovedResponse\"_\n\x0fRegisterRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04port\x18\x02 \x01(\r\x12\x0c\n\x04mcpu\x18\x03 \x01(\r\x12\x0e\n\x06memory\x18\x04 \x01(\x04\x12\x14\n\x0c\x64\x65\x66\x61ult_mcpu\x18\x05 \x01(\r\"\x12\n\x10RegisterResponse\".\n\x10HeartbeatRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0e\n\x06memory\x18\x02 \x01(\x04\"\x13\n\x11HeartbeatResponse\"\x1f\n\x11\x44isconnectRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\x14\n\x12\x44isconnectResponse2\xa1\x03\n\nWorkerPlug\x12s\n\x10\x43ontainerRemoved\x12..syncano.codebox.lb.v1.ContainerRemovedRequest\x1a/.syncano.codebox.lb.v1.ContainerRemovedResponse\x12[\n\x08Register\x12&.syncano.codebox.lb.v1.RegisterRequest\x1a\'.syncano.codebox.lb.v1.RegisterResponse\x12^\n\tHeartbeat\x12\'.syncano.codebox.lb.v1.HeartbeatRequest\x1a(.syncano.codebox.lb.v1.HeartbeatResponse\x12\x61\n\nDisconnect\x12(.syncano.codebox.lb.v1.DisconnectRequest\x1a).syncano.codebox.lb.v1.DisconnectResponseB@Z>github.com/Syncano/syncanoapis/gen/go/syncano/codebox/lb/v1;lbb\x06proto3'
)




_CONTAINERREMOVEDREQUEST = _descriptor.Descriptor(
  name='ContainerRemovedRequest',
  full_name='syncano.codebox.lb.v1.ContainerRemovedRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='syncano.codebox.lb.v1.ContainerRemovedRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='container_id', full_name='syncano.codebox.lb.v1.ContainerRemovedRequest.container_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source_hash', full_name='syncano.codebox.lb.v1.ContainerRemovedRequest.source_hash', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='environment', full_name='syncano.codebox.lb.v1.ContainerRemovedRequest.environment', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='syncano.codebox.lb.v1.ContainerRemovedRequest.user_id', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=65,
  serialized_end=183,
)


_CONTAINERREMOVEDRESPONSE = _descriptor.Descriptor(
  name='ContainerRemovedResponse',
  full_name='syncano.codebox.lb.v1.ContainerRemovedResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=185,
  serialized_end=211,
)


_REGISTERREQUEST = _descriptor.Descriptor(
  name='RegisterRequest',
  full_name='syncano.codebox.lb.v1.RegisterRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='syncano.codebox.lb.v1.RegisterRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='port', full_name='syncano.codebox.lb.v1.RegisterRequest.port', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mcpu', full_name='syncano.codebox.lb.v1.RegisterRequest.mcpu', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='memory', full_name='syncano.codebox.lb.v1.RegisterRequest.memory', index=3,
      number=4, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_mcpu', full_name='syncano.codebox.lb.v1.RegisterRequest.default_mcpu', index=4,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=213,
  serialized_end=308,
)


_REGISTERRESPONSE = _descriptor.Descriptor(
  name='RegisterResponse',
  full_name='syncano.codebox.lb.v1.RegisterResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=310,
  serialized_end=328,
)


_HEARTBEATREQUEST = _descriptor.Descriptor(
  name='HeartbeatRequest',
  full_name='syncano.codebox.lb.v1.HeartbeatRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='syncano.codebox.lb.v1.HeartbeatRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='memory', full_name='syncano.codebox.lb.v1.HeartbeatRequest.memory', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=330,
  serialized_end=376,
)


_HEARTBEATRESPONSE = _descriptor.Descriptor(
  name='HeartbeatResponse',
  full_name='syncano.codebox.lb.v1.HeartbeatResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=378,
  serialized_end=397,
)


_DISCONNECTREQUEST = _descriptor.Descriptor(
  name='DisconnectRequest',
  full_name='syncano.codebox.lb.v1.DisconnectRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='syncano.codebox.lb.v1.DisconnectRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=399,
  serialized_end=430,
)


_DISCONNECTRESPONSE = _descriptor.Descriptor(
  name='DisconnectResponse',
  full_name='syncano.codebox.lb.v1.DisconnectResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=432,
  serialized_end=452,
)

DESCRIPTOR.message_types_by_name['ContainerRemovedRequest'] = _CONTAINERREMOVEDREQUEST
DESCRIPTOR.message_types_by_name['ContainerRemovedResponse'] = _CONTAINERREMOVEDRESPONSE
DESCRIPTOR.message_types_by_name['RegisterRequest'] = _REGISTERREQUEST
DESCRIPTOR.message_types_by_name['RegisterResponse'] = _REGISTERRESPONSE
DESCRIPTOR.message_types_by_name['HeartbeatRequest'] = _HEARTBEATREQUEST
DESCRIPTOR.message_types_by_name['HeartbeatResponse'] = _HEARTBEATRESPONSE
DESCRIPTOR.message_types_by_name['DisconnectRequest'] = _DISCONNECTREQUEST
DESCRIPTOR.message_types_by_name['DisconnectResponse'] = _DISCONNECTRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ContainerRemovedRequest = _reflection.GeneratedProtocolMessageType('ContainerRemovedRequest', (_message.Message,), {
  'DESCRIPTOR' : _CONTAINERREMOVEDREQUEST,
  '__module__' : 'syncano.codebox.lb.v1.workerplug_pb2'
  # @@protoc_insertion_point(class_scope:syncano.codebox.lb.v1.ContainerRemovedRequest)
  })
_sym_db.RegisterMessage(ContainerRemovedRequest)

ContainerRemovedResponse = _reflection.GeneratedProtocolMessageType('ContainerRemovedResponse', (_message.Message,), {
  'DESCRIPTOR' : _CONTAINERREMOVEDRESPONSE,
  '__module__' : 'syncano.codebox.lb.v1.workerplug_pb2'
  # @@protoc_insertion_point(class_scope:syncano.codebox.lb.v1.ContainerRemovedResponse)
  })
_sym_db.RegisterMessage(ContainerRemovedResponse)

RegisterRequest = _reflection.GeneratedProtocolMessageType('RegisterRequest', (_message.Message,), {
  'DESCRIPTOR' : _REGISTERREQUEST,
  '__module__' : 'syncano.codebox.lb.v1.workerplug_pb2'
  # @@protoc_insertion_point(class_scope:syncano.codebox.lb.v1.RegisterRequest)
  })
_sym_db.RegisterMessage(RegisterRequest)

RegisterResponse = _reflection.GeneratedProtocolMessageType('RegisterResponse', (_message.Message,), {
  'DESCRIPTOR' : _REGISTERRESPONSE,
  '__module__' : 'syncano.codebox.lb.v1.workerplug_pb2'
  # @@protoc_insertion_point(class_scope:syncano.codebox.lb.v1.RegisterResponse)
  })
_sym_db.RegisterMessage(RegisterResponse)

HeartbeatRequest = _reflection.GeneratedProtocolMessageType('HeartbeatRequest', (_message.Message,), {
  'DESCRIPTOR' : _HEARTBEATREQUEST,
  '__module__' : 'syncano.codebox.lb.v1.workerplug_pb2'
  # @@protoc_insertion_point(class_scope:syncano.codebox.lb.v1.HeartbeatRequest)
  })
_sym_db.RegisterMessage(HeartbeatRequest)

HeartbeatResponse = _reflection.GeneratedProtocolMessageType('HeartbeatResponse', (_message.Message,), {
  'DESCRIPTOR' : _HEARTBEATRESPONSE,
  '__module__' : 'syncano.codebox.lb.v1.workerplug_pb2'
  # @@protoc_insertion_point(class_scope:syncano.codebox.lb.v1.HeartbeatResponse)
  })
_sym_db.RegisterMessage(HeartbeatResponse)

DisconnectRequest = _reflection.GeneratedProtocolMessageType('DisconnectRequest', (_message.Message,), {
  'DESCRIPTOR' : _DISCONNECTREQUEST,
  '__module__' : 'syncano.codebox.lb.v1.workerplug_pb2'
  # @@protoc_insertion_point(class_scope:syncano.codebox.lb.v1.DisconnectRequest)
  })
_sym_db.RegisterMessage(DisconnectRequest)

DisconnectResponse = _reflection.GeneratedProtocolMessageType('DisconnectResponse', (_message.Message,), {
  'DESCRIPTOR' : _DISCONNECTRESPONSE,
  '__module__' : 'syncano.codebox.lb.v1.workerplug_pb2'
  # @@protoc_insertion_point(class_scope:syncano.codebox.lb.v1.DisconnectResponse)
  })
_sym_db.RegisterMessage(DisconnectResponse)


DESCRIPTOR._options = None

_WORKERPLUG = _descriptor.ServiceDescriptor(
  name='WorkerPlug',
  full_name='syncano.codebox.lb.v1.WorkerPlug',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=455,
  serialized_end=872,
  methods=[
  _descriptor.MethodDescriptor(
    name='ContainerRemoved',
    full_name='syncano.codebox.lb.v1.WorkerPlug.ContainerRemoved',
    index=0,
    containing_service=None,
    input_type=_CONTAINERREMOVEDREQUEST,
    output_type=_CONTAINERREMOVEDRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Register',
    full_name='syncano.codebox.lb.v1.WorkerPlug.Register',
    index=1,
    containing_service=None,
    input_type=_REGISTERREQUEST,
    output_type=_REGISTERRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Heartbeat',
    full_name='syncano.codebox.lb.v1.WorkerPlug.Heartbeat',
    index=2,
    containing_service=None,
    input_type=_HEARTBEATREQUEST,
    output_type=_HEARTBEATRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Disconnect',
    full_name='syncano.codebox.lb.v1.WorkerPlug.Disconnect',
    index=3,
    containing_service=None,
    input_type=_DISCONNECTREQUEST,
    output_type=_DISCONNECTRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_WORKERPLUG)

DESCRIPTOR.services_by_name['WorkerPlug'] = _WORKERPLUG

# @@protoc_insertion_point(module_scope)
