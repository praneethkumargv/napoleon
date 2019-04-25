# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: napolet.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='napolet.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\rnapolet.proto\"\x07\n\x05\x44ummy\"\x80\x01\n\x0f\x43reateVMRequest\x12\x0e\n\x06VMName\x18\x01 \x01(\t\x12\r\n\x05vcpus\x18\x02 \x01(\r\x12\x0e\n\x06memory\x18\x03 \x01(\r\x12\x0f\n\x07storage\x18\x04 \x01(\r\x12\x11\n\tImageName\x18\x05 \x01(\t\x12\x0c\n\x04VMId\x18\x06 \x01(\t\x12\x0c\n\x04PMId\x18\x07 \x01(\t\"$\n\x10\x43reateVMResponse\x12\x10\n\x08\x61\x63\x63\x65pted\x18\x01 \x01(\x08\"\x1f\n\x0f\x44\x65leteVMRequest\x12\x0c\n\x04VMId\x18\x01 \x01(\t\"$\n\x10\x44\x65leteVMResponse\x12\x10\n\x08\x61\x63\x63\x65pted\x18\x01 \x01(\x08\"T\n\x06VMStat\x12\x0c\n\x04VMId\x18\x01 \x01(\t\x12\r\n\x05state\x18\x02 \x01(\t\x12\x14\n\x0cPredictedCpu\x18\x03 \x01(\r\x12\x17\n\x0fPredictedMemory\x18\x04 \x01(\r\"d\n\x06PMStat\x12\x0c\n\x04PMId\x18\x01 \x01(\t\x12\x13\n\x0bTotalMemory\x18\x02 \x01(\r\x12\x10\n\x08TotalCpu\x18\x03 \x01(\r\x12\x10\n\x08SlackCpu\x18\x04 \x01(\r\x12\x13\n\x0bSlackMemory\x18\x05 \x01(\r\"1\n\x04Stat\x12\x13\n\x02PM\x18\x01 \x01(\x0b\x32\x07.PMStat\x12\x14\n\x03VMS\x18\x02 \x03(\x0b\x32\x07.VMStat2\x88\x01\n\x04Ping\x12\x31\n\x08\x43reateVM\x12\x10.CreateVMRequest\x1a\x11.CreateVMResponse\"\x00\x12\x31\n\x08\x44\x65leteVM\x12\x10.DeleteVMRequest\x1a\x11.DeleteVMResponse\"\x00\x12\x1a\n\x07GetStat\x12\x06.Dummy\x1a\x05.Stat\"\x00\x62\x06proto3')
)




_DUMMY = _descriptor.Descriptor(
  name='Dummy',
  full_name='Dummy',
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
  serialized_start=17,
  serialized_end=24,
)


_CREATEVMREQUEST = _descriptor.Descriptor(
  name='CreateVMRequest',
  full_name='CreateVMRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='VMName', full_name='CreateVMRequest.VMName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vcpus', full_name='CreateVMRequest.vcpus', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='memory', full_name='CreateVMRequest.memory', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage', full_name='CreateVMRequest.storage', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ImageName', full_name='CreateVMRequest.ImageName', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='VMId', full_name='CreateVMRequest.VMId', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='PMId', full_name='CreateVMRequest.PMId', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=27,
  serialized_end=155,
)


_CREATEVMRESPONSE = _descriptor.Descriptor(
  name='CreateVMResponse',
  full_name='CreateVMResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='accepted', full_name='CreateVMResponse.accepted', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=157,
  serialized_end=193,
)


_DELETEVMREQUEST = _descriptor.Descriptor(
  name='DeleteVMRequest',
  full_name='DeleteVMRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='VMId', full_name='DeleteVMRequest.VMId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=195,
  serialized_end=226,
)


_DELETEVMRESPONSE = _descriptor.Descriptor(
  name='DeleteVMResponse',
  full_name='DeleteVMResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='accepted', full_name='DeleteVMResponse.accepted', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=228,
  serialized_end=264,
)


_VMSTAT = _descriptor.Descriptor(
  name='VMStat',
  full_name='VMStat',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='VMId', full_name='VMStat.VMId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='state', full_name='VMStat.state', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='PredictedCpu', full_name='VMStat.PredictedCpu', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='PredictedMemory', full_name='VMStat.PredictedMemory', index=3,
      number=4, type=13, cpp_type=3, label=1,
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
  serialized_start=266,
  serialized_end=350,
)


_PMSTAT = _descriptor.Descriptor(
  name='PMStat',
  full_name='PMStat',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='PMId', full_name='PMStat.PMId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='TotalMemory', full_name='PMStat.TotalMemory', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='TotalCpu', full_name='PMStat.TotalCpu', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='SlackCpu', full_name='PMStat.SlackCpu', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='SlackMemory', full_name='PMStat.SlackMemory', index=4,
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
  serialized_start=352,
  serialized_end=452,
)


_STAT = _descriptor.Descriptor(
  name='Stat',
  full_name='Stat',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='PM', full_name='Stat.PM', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='VMS', full_name='Stat.VMS', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=454,
  serialized_end=503,
)

_STAT.fields_by_name['PM'].message_type = _PMSTAT
_STAT.fields_by_name['VMS'].message_type = _VMSTAT
DESCRIPTOR.message_types_by_name['Dummy'] = _DUMMY
DESCRIPTOR.message_types_by_name['CreateVMRequest'] = _CREATEVMREQUEST
DESCRIPTOR.message_types_by_name['CreateVMResponse'] = _CREATEVMRESPONSE
DESCRIPTOR.message_types_by_name['DeleteVMRequest'] = _DELETEVMREQUEST
DESCRIPTOR.message_types_by_name['DeleteVMResponse'] = _DELETEVMRESPONSE
DESCRIPTOR.message_types_by_name['VMStat'] = _VMSTAT
DESCRIPTOR.message_types_by_name['PMStat'] = _PMSTAT
DESCRIPTOR.message_types_by_name['Stat'] = _STAT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Dummy = _reflection.GeneratedProtocolMessageType('Dummy', (_message.Message,), dict(
  DESCRIPTOR = _DUMMY,
  __module__ = 'napolet_pb2'
  # @@protoc_insertion_point(class_scope:Dummy)
  ))
_sym_db.RegisterMessage(Dummy)

CreateVMRequest = _reflection.GeneratedProtocolMessageType('CreateVMRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATEVMREQUEST,
  __module__ = 'napolet_pb2'
  # @@protoc_insertion_point(class_scope:CreateVMRequest)
  ))
_sym_db.RegisterMessage(CreateVMRequest)

CreateVMResponse = _reflection.GeneratedProtocolMessageType('CreateVMResponse', (_message.Message,), dict(
  DESCRIPTOR = _CREATEVMRESPONSE,
  __module__ = 'napolet_pb2'
  # @@protoc_insertion_point(class_scope:CreateVMResponse)
  ))
_sym_db.RegisterMessage(CreateVMResponse)

DeleteVMRequest = _reflection.GeneratedProtocolMessageType('DeleteVMRequest', (_message.Message,), dict(
  DESCRIPTOR = _DELETEVMREQUEST,
  __module__ = 'napolet_pb2'
  # @@protoc_insertion_point(class_scope:DeleteVMRequest)
  ))
_sym_db.RegisterMessage(DeleteVMRequest)

DeleteVMResponse = _reflection.GeneratedProtocolMessageType('DeleteVMResponse', (_message.Message,), dict(
  DESCRIPTOR = _DELETEVMRESPONSE,
  __module__ = 'napolet_pb2'
  # @@protoc_insertion_point(class_scope:DeleteVMResponse)
  ))
_sym_db.RegisterMessage(DeleteVMResponse)

VMStat = _reflection.GeneratedProtocolMessageType('VMStat', (_message.Message,), dict(
  DESCRIPTOR = _VMSTAT,
  __module__ = 'napolet_pb2'
  # @@protoc_insertion_point(class_scope:VMStat)
  ))
_sym_db.RegisterMessage(VMStat)

PMStat = _reflection.GeneratedProtocolMessageType('PMStat', (_message.Message,), dict(
  DESCRIPTOR = _PMSTAT,
  __module__ = 'napolet_pb2'
  # @@protoc_insertion_point(class_scope:PMStat)
  ))
_sym_db.RegisterMessage(PMStat)

Stat = _reflection.GeneratedProtocolMessageType('Stat', (_message.Message,), dict(
  DESCRIPTOR = _STAT,
  __module__ = 'napolet_pb2'
  # @@protoc_insertion_point(class_scope:Stat)
  ))
_sym_db.RegisterMessage(Stat)



_PING = _descriptor.ServiceDescriptor(
  name='Ping',
  full_name='Ping',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=506,
  serialized_end=642,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateVM',
    full_name='Ping.CreateVM',
    index=0,
    containing_service=None,
    input_type=_CREATEVMREQUEST,
    output_type=_CREATEVMRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteVM',
    full_name='Ping.DeleteVM',
    index=1,
    containing_service=None,
    input_type=_DELETEVMREQUEST,
    output_type=_DELETEVMRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetStat',
    full_name='Ping.GetStat',
    index=2,
    containing_service=None,
    input_type=_DUMMY,
    output_type=_STAT,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_PING)

DESCRIPTOR.services_by_name['Ping'] = _PING

# @@protoc_insertion_point(module_scope)
