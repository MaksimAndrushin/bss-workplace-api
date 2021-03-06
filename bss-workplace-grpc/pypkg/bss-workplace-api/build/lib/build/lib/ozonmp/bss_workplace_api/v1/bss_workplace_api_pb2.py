# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ozonmp/bss_workplace_api/v1/bss_workplace_api.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='ozonmp/bss_workplace_api/v1/bss_workplace_api.proto',
  package='ozonmp.bss_workplace_api.v1',
  syntax='proto3',
  serialized_options=_b('ZKgithub.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api;bss_workplace_api'),
  serialized_pb=_b('\n3ozonmp/bss_workplace_api/v1/bss_workplace_api.proto\x12\x1bozonmp.bss_workplace_api.v1\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"c\n\tWorkplace\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x10\n\x03\x66oo\x18\x02 \x01(\tR\x03\x66oo\x12\x34\n\x07\x63reated\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07\x63reated\",\n\x18\x43reateWorkplaceV1Request\x12\x10\n\x03\x66oo\x18\x01 \x01(\tR\x03\x66oo\">\n\x19\x43reateWorkplaceV1Response\x12!\n\x0cworkplace_id\x18\x01 \x01(\x04R\x0bworkplaceId\"H\n\x1a\x44\x65scribeWorkplaceV1Request\x12*\n\x0cworkplace_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x0bworkplaceId\"[\n\x1b\x44\x65scribeWorkplaceV1Response\x12<\n\x05value\x18\x01 \x01(\x0b\x32&.ozonmp.bss_workplace_api.v1.WorkplaceR\x05value\"\x19\n\x17ListWorkplacesV1Request\"X\n\x18ListWorkplacesV1Response\x12<\n\x05items\x18\x01 \x03(\x0b\x32&.ozonmp.bss_workplace_api.v1.WorkplaceR\x05items\"F\n\x18RemoveWorkplaceV1Request\x12*\n\x0cworkplace_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x0bworkplaceId\"1\n\x19RemoveWorkplaceV1Response\x12\x14\n\x05\x66ound\x18\x01 \x01(\x08R\x05\x66ound2\xb2\x05\n\x16\x42ssWorkplaceApiService\x12\x9f\x01\n\x11\x43reateWorkplaceV1\x12\x35.ozonmp.bss_workplace_api.v1.CreateWorkplaceV1Request\x1a\x36.ozonmp.bss_workplace_api.v1.CreateWorkplaceV1Response\"\x1b\x82\xd3\xe4\x93\x02\x15\"\x0e/v1/workplaces:\x03\x66oo\x12\xaf\x01\n\x13\x44\x65scribeWorkplaceV1\x12\x37.ozonmp.bss_workplace_api.v1.DescribeWorkplaceV1Request\x1a\x38.ozonmp.bss_workplace_api.v1.DescribeWorkplaceV1Response\"%\x82\xd3\xe4\x93\x02\x1f\x12\x1d/v1/workplaces/{workplace_id}\x12\x97\x01\n\x10ListWorkplacesV1\x12\x34.ozonmp.bss_workplace_api.v1.ListWorkplacesV1Request\x1a\x35.ozonmp.bss_workplace_api.v1.ListWorkplacesV1Response\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/v1/workplaces\x12\xa9\x01\n\x11RemoveWorkplaceV1\x12\x35.ozonmp.bss_workplace_api.v1.RemoveWorkplaceV1Request\x1a\x36.ozonmp.bss_workplace_api.v1.RemoveWorkplaceV1Response\"%\x82\xd3\xe4\x93\x02\x1f*\x1d/v1/workplaces/{workplace_id}BMZKgithub.com/ozonmp/bss-workplace-api/pkg/bss-workplace-api;bss_workplace_apib\x06proto3')
  ,
  dependencies=[validate_dot_validate__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_WORKPLACE = _descriptor.Descriptor(
  name='Workplace',
  full_name='ozonmp.bss_workplace_api.v1.Workplace',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='ozonmp.bss_workplace_api.v1.Workplace.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='id', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='foo', full_name='ozonmp.bss_workplace_api.v1.Workplace.foo', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='foo', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created', full_name='ozonmp.bss_workplace_api.v1.Workplace.created', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='created', file=DESCRIPTOR),
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
  serialized_start=172,
  serialized_end=271,
)


_CREATEWORKPLACEV1REQUEST = _descriptor.Descriptor(
  name='CreateWorkplaceV1Request',
  full_name='ozonmp.bss_workplace_api.v1.CreateWorkplaceV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='foo', full_name='ozonmp.bss_workplace_api.v1.CreateWorkplaceV1Request.foo', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='foo', file=DESCRIPTOR),
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
  serialized_start=273,
  serialized_end=317,
)


_CREATEWORKPLACEV1RESPONSE = _descriptor.Descriptor(
  name='CreateWorkplaceV1Response',
  full_name='ozonmp.bss_workplace_api.v1.CreateWorkplaceV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='workplace_id', full_name='ozonmp.bss_workplace_api.v1.CreateWorkplaceV1Response.workplace_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='workplaceId', file=DESCRIPTOR),
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
  serialized_start=319,
  serialized_end=381,
)


_DESCRIBEWORKPLACEV1REQUEST = _descriptor.Descriptor(
  name='DescribeWorkplaceV1Request',
  full_name='ozonmp.bss_workplace_api.v1.DescribeWorkplaceV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='workplace_id', full_name='ozonmp.bss_workplace_api.v1.DescribeWorkplaceV1Request.workplace_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\0042\002 \000'), json_name='workplaceId', file=DESCRIPTOR),
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
  serialized_start=383,
  serialized_end=455,
)


_DESCRIBEWORKPLACEV1RESPONSE = _descriptor.Descriptor(
  name='DescribeWorkplaceV1Response',
  full_name='ozonmp.bss_workplace_api.v1.DescribeWorkplaceV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='ozonmp.bss_workplace_api.v1.DescribeWorkplaceV1Response.value', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR),
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
  serialized_start=457,
  serialized_end=548,
)


_LISTWORKPLACESV1REQUEST = _descriptor.Descriptor(
  name='ListWorkplacesV1Request',
  full_name='ozonmp.bss_workplace_api.v1.ListWorkplacesV1Request',
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
  serialized_start=550,
  serialized_end=575,
)


_LISTWORKPLACESV1RESPONSE = _descriptor.Descriptor(
  name='ListWorkplacesV1Response',
  full_name='ozonmp.bss_workplace_api.v1.ListWorkplacesV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='ozonmp.bss_workplace_api.v1.ListWorkplacesV1Response.items', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='items', file=DESCRIPTOR),
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
  serialized_start=577,
  serialized_end=665,
)


_REMOVEWORKPLACEV1REQUEST = _descriptor.Descriptor(
  name='RemoveWorkplaceV1Request',
  full_name='ozonmp.bss_workplace_api.v1.RemoveWorkplaceV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='workplace_id', full_name='ozonmp.bss_workplace_api.v1.RemoveWorkplaceV1Request.workplace_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\0042\002 \000'), json_name='workplaceId', file=DESCRIPTOR),
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
  serialized_start=667,
  serialized_end=737,
)


_REMOVEWORKPLACEV1RESPONSE = _descriptor.Descriptor(
  name='RemoveWorkplaceV1Response',
  full_name='ozonmp.bss_workplace_api.v1.RemoveWorkplaceV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='found', full_name='ozonmp.bss_workplace_api.v1.RemoveWorkplaceV1Response.found', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='found', file=DESCRIPTOR),
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
  serialized_start=739,
  serialized_end=788,
)

_WORKPLACE.fields_by_name['created'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_DESCRIBEWORKPLACEV1RESPONSE.fields_by_name['value'].message_type = _WORKPLACE
_LISTWORKPLACESV1RESPONSE.fields_by_name['items'].message_type = _WORKPLACE
DESCRIPTOR.message_types_by_name['Workplace'] = _WORKPLACE
DESCRIPTOR.message_types_by_name['CreateWorkplaceV1Request'] = _CREATEWORKPLACEV1REQUEST
DESCRIPTOR.message_types_by_name['CreateWorkplaceV1Response'] = _CREATEWORKPLACEV1RESPONSE
DESCRIPTOR.message_types_by_name['DescribeWorkplaceV1Request'] = _DESCRIBEWORKPLACEV1REQUEST
DESCRIPTOR.message_types_by_name['DescribeWorkplaceV1Response'] = _DESCRIBEWORKPLACEV1RESPONSE
DESCRIPTOR.message_types_by_name['ListWorkplacesV1Request'] = _LISTWORKPLACESV1REQUEST
DESCRIPTOR.message_types_by_name['ListWorkplacesV1Response'] = _LISTWORKPLACESV1RESPONSE
DESCRIPTOR.message_types_by_name['RemoveWorkplaceV1Request'] = _REMOVEWORKPLACEV1REQUEST
DESCRIPTOR.message_types_by_name['RemoveWorkplaceV1Response'] = _REMOVEWORKPLACEV1RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Workplace = _reflection.GeneratedProtocolMessageType('Workplace', (_message.Message,), dict(
  DESCRIPTOR = _WORKPLACE,
  __module__ = 'ozonmp.bss_workplace_api.v1.bss_workplace_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_workplace_api.v1.Workplace)
  ))
_sym_db.RegisterMessage(Workplace)

CreateWorkplaceV1Request = _reflection.GeneratedProtocolMessageType('CreateWorkplaceV1Request', (_message.Message,), dict(
  DESCRIPTOR = _CREATEWORKPLACEV1REQUEST,
  __module__ = 'ozonmp.bss_workplace_api.v1.bss_workplace_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_workplace_api.v1.CreateWorkplaceV1Request)
  ))
_sym_db.RegisterMessage(CreateWorkplaceV1Request)

CreateWorkplaceV1Response = _reflection.GeneratedProtocolMessageType('CreateWorkplaceV1Response', (_message.Message,), dict(
  DESCRIPTOR = _CREATEWORKPLACEV1RESPONSE,
  __module__ = 'ozonmp.bss_workplace_api.v1.bss_workplace_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_workplace_api.v1.CreateWorkplaceV1Response)
  ))
_sym_db.RegisterMessage(CreateWorkplaceV1Response)

DescribeWorkplaceV1Request = _reflection.GeneratedProtocolMessageType('DescribeWorkplaceV1Request', (_message.Message,), dict(
  DESCRIPTOR = _DESCRIBEWORKPLACEV1REQUEST,
  __module__ = 'ozonmp.bss_workplace_api.v1.bss_workplace_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_workplace_api.v1.DescribeWorkplaceV1Request)
  ))
_sym_db.RegisterMessage(DescribeWorkplaceV1Request)

DescribeWorkplaceV1Response = _reflection.GeneratedProtocolMessageType('DescribeWorkplaceV1Response', (_message.Message,), dict(
  DESCRIPTOR = _DESCRIBEWORKPLACEV1RESPONSE,
  __module__ = 'ozonmp.bss_workplace_api.v1.bss_workplace_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_workplace_api.v1.DescribeWorkplaceV1Response)
  ))
_sym_db.RegisterMessage(DescribeWorkplaceV1Response)

ListWorkplacesV1Request = _reflection.GeneratedProtocolMessageType('ListWorkplacesV1Request', (_message.Message,), dict(
  DESCRIPTOR = _LISTWORKPLACESV1REQUEST,
  __module__ = 'ozonmp.bss_workplace_api.v1.bss_workplace_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_workplace_api.v1.ListWorkplacesV1Request)
  ))
_sym_db.RegisterMessage(ListWorkplacesV1Request)

ListWorkplacesV1Response = _reflection.GeneratedProtocolMessageType('ListWorkplacesV1Response', (_message.Message,), dict(
  DESCRIPTOR = _LISTWORKPLACESV1RESPONSE,
  __module__ = 'ozonmp.bss_workplace_api.v1.bss_workplace_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_workplace_api.v1.ListWorkplacesV1Response)
  ))
_sym_db.RegisterMessage(ListWorkplacesV1Response)

RemoveWorkplaceV1Request = _reflection.GeneratedProtocolMessageType('RemoveWorkplaceV1Request', (_message.Message,), dict(
  DESCRIPTOR = _REMOVEWORKPLACEV1REQUEST,
  __module__ = 'ozonmp.bss_workplace_api.v1.bss_workplace_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_workplace_api.v1.RemoveWorkplaceV1Request)
  ))
_sym_db.RegisterMessage(RemoveWorkplaceV1Request)

RemoveWorkplaceV1Response = _reflection.GeneratedProtocolMessageType('RemoveWorkplaceV1Response', (_message.Message,), dict(
  DESCRIPTOR = _REMOVEWORKPLACEV1RESPONSE,
  __module__ = 'ozonmp.bss_workplace_api.v1.bss_workplace_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.bss_workplace_api.v1.RemoveWorkplaceV1Response)
  ))
_sym_db.RegisterMessage(RemoveWorkplaceV1Response)


DESCRIPTOR._options = None
_DESCRIBEWORKPLACEV1REQUEST.fields_by_name['workplace_id']._options = None
_REMOVEWORKPLACEV1REQUEST.fields_by_name['workplace_id']._options = None

_BSSWORKPLACEAPISERVICE = _descriptor.ServiceDescriptor(
  name='BssWorkplaceApiService',
  full_name='ozonmp.bss_workplace_api.v1.BssWorkplaceApiService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=791,
  serialized_end=1481,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateWorkplaceV1',
    full_name='ozonmp.bss_workplace_api.v1.BssWorkplaceApiService.CreateWorkplaceV1',
    index=0,
    containing_service=None,
    input_type=_CREATEWORKPLACEV1REQUEST,
    output_type=_CREATEWORKPLACEV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\025\"\016/v1/workplaces:\003foo'),
  ),
  _descriptor.MethodDescriptor(
    name='DescribeWorkplaceV1',
    full_name='ozonmp.bss_workplace_api.v1.BssWorkplaceApiService.DescribeWorkplaceV1',
    index=1,
    containing_service=None,
    input_type=_DESCRIBEWORKPLACEV1REQUEST,
    output_type=_DESCRIBEWORKPLACEV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\037\022\035/v1/workplaces/{workplace_id}'),
  ),
  _descriptor.MethodDescriptor(
    name='ListWorkplacesV1',
    full_name='ozonmp.bss_workplace_api.v1.BssWorkplaceApiService.ListWorkplacesV1',
    index=2,
    containing_service=None,
    input_type=_LISTWORKPLACESV1REQUEST,
    output_type=_LISTWORKPLACESV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\020\022\016/v1/workplaces'),
  ),
  _descriptor.MethodDescriptor(
    name='RemoveWorkplaceV1',
    full_name='ozonmp.bss_workplace_api.v1.BssWorkplaceApiService.RemoveWorkplaceV1',
    index=3,
    containing_service=None,
    input_type=_REMOVEWORKPLACEV1REQUEST,
    output_type=_REMOVEWORKPLACEV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\037*\035/v1/workplaces/{workplace_id}'),
  ),
])
_sym_db.RegisterServiceDescriptor(_BSSWORKPLACEAPISERVICE)

DESCRIPTOR.services_by_name['BssWorkplaceApiService'] = _BSSWORKPLACEAPISERVICE

# @@protoc_insertion_point(module_scope)
