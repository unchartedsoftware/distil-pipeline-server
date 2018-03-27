// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pipeline.proto

package pipeline

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/protoc-gen-go/descriptor"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContainerArgument struct {
	// Data reference.
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *ContainerArgument) Reset()                    { *m = ContainerArgument{} }
func (m *ContainerArgument) String() string            { return proto.CompactTextString(m) }
func (*ContainerArgument) ProtoMessage()               {}
func (*ContainerArgument) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ContainerArgument) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PrimitiveArgument struct {
	// 0-based index from steps identifying a primitive to pass in.
	Primitive string `protobuf:"bytes,1,opt,name=primitive" json:"primitive,omitempty"`
}

func (m *PrimitiveArgument) Reset()                    { *m = PrimitiveArgument{} }
func (m *PrimitiveArgument) String() string            { return proto.CompactTextString(m) }
func (*PrimitiveArgument) ProtoMessage()               {}
func (*PrimitiveArgument) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PrimitiveArgument) GetPrimitive() string {
	if m != nil {
		return m.Primitive
	}
	return ""
}

type DataArgument struct {
	// Data reference.
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *DataArgument) Reset()                    { *m = DataArgument{} }
func (m *DataArgument) String() string            { return proto.CompactTextString(m) }
func (*DataArgument) ProtoMessage()               {}
func (*DataArgument) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *DataArgument) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PrimitiveStepArgument struct {
	// Types that are valid to be assigned to Argument:
	//	*PrimitiveStepArgument_Container
	//	*PrimitiveStepArgument_Primitive
	//	*PrimitiveStepArgument_Data
	Argument isPrimitiveStepArgument_Argument `protobuf_oneof:"argument"`
}

func (m *PrimitiveStepArgument) Reset()                    { *m = PrimitiveStepArgument{} }
func (m *PrimitiveStepArgument) String() string            { return proto.CompactTextString(m) }
func (*PrimitiveStepArgument) ProtoMessage()               {}
func (*PrimitiveStepArgument) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type isPrimitiveStepArgument_Argument interface {
	isPrimitiveStepArgument_Argument()
}

type PrimitiveStepArgument_Container struct {
	Container *ContainerArgument `protobuf:"bytes,1,opt,name=container,oneof"`
}
type PrimitiveStepArgument_Primitive struct {
	Primitive *PrimitiveArgument `protobuf:"bytes,2,opt,name=primitive,oneof"`
}
type PrimitiveStepArgument_Data struct {
	Data *DataArgument `protobuf:"bytes,3,opt,name=data,oneof"`
}

func (*PrimitiveStepArgument_Container) isPrimitiveStepArgument_Argument() {}
func (*PrimitiveStepArgument_Primitive) isPrimitiveStepArgument_Argument() {}
func (*PrimitiveStepArgument_Data) isPrimitiveStepArgument_Argument()      {}

func (m *PrimitiveStepArgument) GetArgument() isPrimitiveStepArgument_Argument {
	if m != nil {
		return m.Argument
	}
	return nil
}

func (m *PrimitiveStepArgument) GetContainer() *ContainerArgument {
	if x, ok := m.GetArgument().(*PrimitiveStepArgument_Container); ok {
		return x.Container
	}
	return nil
}

func (m *PrimitiveStepArgument) GetPrimitive() *PrimitiveArgument {
	if x, ok := m.GetArgument().(*PrimitiveStepArgument_Primitive); ok {
		return x.Primitive
	}
	return nil
}

func (m *PrimitiveStepArgument) GetData() *DataArgument {
	if x, ok := m.GetArgument().(*PrimitiveStepArgument_Data); ok {
		return x.Data
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PrimitiveStepArgument) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PrimitiveStepArgument_OneofMarshaler, _PrimitiveStepArgument_OneofUnmarshaler, _PrimitiveStepArgument_OneofSizer, []interface{}{
		(*PrimitiveStepArgument_Container)(nil),
		(*PrimitiveStepArgument_Primitive)(nil),
		(*PrimitiveStepArgument_Data)(nil),
	}
}

func _PrimitiveStepArgument_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PrimitiveStepArgument)
	// argument
	switch x := m.Argument.(type) {
	case *PrimitiveStepArgument_Container:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Container); err != nil {
			return err
		}
	case *PrimitiveStepArgument_Primitive:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Primitive); err != nil {
			return err
		}
	case *PrimitiveStepArgument_Data:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Data); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PrimitiveStepArgument.Argument has unexpected type %T", x)
	}
	return nil
}

func _PrimitiveStepArgument_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PrimitiveStepArgument)
	switch tag {
	case 1: // argument.container
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContainerArgument)
		err := b.DecodeMessage(msg)
		m.Argument = &PrimitiveStepArgument_Container{msg}
		return true, err
	case 2: // argument.primitive
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrimitiveArgument)
		err := b.DecodeMessage(msg)
		m.Argument = &PrimitiveStepArgument_Primitive{msg}
		return true, err
	case 3: // argument.data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DataArgument)
		err := b.DecodeMessage(msg)
		m.Argument = &PrimitiveStepArgument_Data{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PrimitiveStepArgument_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PrimitiveStepArgument)
	// argument
	switch x := m.Argument.(type) {
	case *PrimitiveStepArgument_Container:
		s := proto.Size(x.Container)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PrimitiveStepArgument_Primitive:
		s := proto.Size(x.Primitive)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PrimitiveStepArgument_Data:
		s := proto.Size(x.Data)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StepInput struct {
	// Data reference.
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *StepInput) Reset()                    { *m = StepInput{} }
func (m *StepInput) String() string            { return proto.CompactTextString(m) }
func (*StepInput) ProtoMessage()               {}
func (*StepInput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *StepInput) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type StepOutput struct {
	// Name which becomes part of the data reference.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StepOutput) Reset()                    { *m = StepOutput{} }
func (m *StepOutput) String() string            { return proto.CompactTextString(m) }
func (*StepOutput) ProtoMessage()               {}
func (*StepOutput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *StepOutput) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PipelineSource struct {
	// String representing name of the author, team.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PipelineSource) Reset()                    { *m = PipelineSource{} }
func (m *PipelineSource) String() string            { return proto.CompactTextString(m) }
func (*PipelineSource) ProtoMessage()               {}
func (*PipelineSource) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *PipelineSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// User associated with the creation of the template/pipeline, or selection of a primitive.
type PipelineDescriptionUser struct {
	// A UUID of the user. It does not have to map to any real ID, just that it is possible
	// to connect mutliple pipelines/templates by the same user together, if necessary.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Textual description of what user did to create the template/pipeline, or select a primitive.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *PipelineDescriptionUser) Reset()                    { *m = PipelineDescriptionUser{} }
func (m *PipelineDescriptionUser) String() string            { return proto.CompactTextString(m) }
func (*PipelineDescriptionUser) ProtoMessage()               {}
func (*PipelineDescriptionUser) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *PipelineDescriptionUser) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PipelineDescriptionUser) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Possible input to the pipeline or template.
type PipelineDescriptionInput struct {
	// Human friendly name of the input.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PipelineDescriptionInput) Reset()                    { *m = PipelineDescriptionInput{} }
func (m *PipelineDescriptionInput) String() string            { return proto.CompactTextString(m) }
func (*PipelineDescriptionInput) ProtoMessage()               {}
func (*PipelineDescriptionInput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *PipelineDescriptionInput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Available output of the pipeline or template.
type PipelineDescriptionOutput struct {
	// Human friendly name of the output.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Data reference, probably of an output of a step.
	Data string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *PipelineDescriptionOutput) Reset()                    { *m = PipelineDescriptionOutput{} }
func (m *PipelineDescriptionOutput) String() string            { return proto.CompactTextString(m) }
func (*PipelineDescriptionOutput) ProtoMessage()               {}
func (*PipelineDescriptionOutput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *PipelineDescriptionOutput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PipelineDescriptionOutput) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type PrimitivePipelineDescriptionStep struct {
	Primitive *Primitive `protobuf:"bytes,1,opt,name=primitive" json:"primitive,omitempty"`
	// Arguments to the primitive. Constructor arguments should not be listed here, because they
	// can be automatically created from other information. All these arguments are listed as kind
	// "PIPELINE" in primitive's metadata.
	Arguments map[string]*PrimitiveStepArgument `protobuf:"bytes,2,rep,name=arguments" json:"arguments,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// List of produce metods providing data. One can reference using data reference these outputs
	// then in arguments (inputs) in other steps or pipeline outputs.
	Outputs []*StepOutput `protobuf:"bytes,3,rep,name=outputs" json:"outputs,omitempty"`
	// Some hyper-parameters are not really tunable and should be fixed as part of template/pipeline.
	// This can be done here. Hyper-parameters listed here cannot be tuned or overridden. Author of a
	// template/pipeline decides which hyper-parameter are which, probably based on their semantic type.
	// TA3 can specify a list of hyper-parameters to fix, and TA2 can add to the list additional
	// hyper-paramaters in found pipelines.
	Hyperparams map[string]*Value `protobuf:"bytes,4,rep,name=hyperparams" json:"hyperparams,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// List of users associated with selection of this primitive/arguments/hyper-parameters. Optional.
	Users []*PipelineDescriptionUser `protobuf:"bytes,5,rep,name=users" json:"users,omitempty"`
}

func (m *PrimitivePipelineDescriptionStep) Reset()         { *m = PrimitivePipelineDescriptionStep{} }
func (m *PrimitivePipelineDescriptionStep) String() string { return proto.CompactTextString(m) }
func (*PrimitivePipelineDescriptionStep) ProtoMessage()    {}
func (*PrimitivePipelineDescriptionStep) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{10}
}

func (m *PrimitivePipelineDescriptionStep) GetPrimitive() *Primitive {
	if m != nil {
		return m.Primitive
	}
	return nil
}

func (m *PrimitivePipelineDescriptionStep) GetArguments() map[string]*PrimitiveStepArgument {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *PrimitivePipelineDescriptionStep) GetOutputs() []*StepOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *PrimitivePipelineDescriptionStep) GetHyperparams() map[string]*Value {
	if m != nil {
		return m.Hyperparams
	}
	return nil
}

func (m *PrimitivePipelineDescriptionStep) GetUsers() []*PipelineDescriptionUser {
	if m != nil {
		return m.Users
	}
	return nil
}

type SubpipelinePipelineDescriptionStep struct {
	// Only "id" field is required in this case to reference another pipeline in the template.
	Pipeline *PipelineDescription `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
	// List of data references, probably of an output of a step or pipeline input,
	// mapped to sub-pipeline's inputs in order.
	Inputs []*StepInput `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty"`
	// List of IDs to be used in data references, mapping sub-pipeline's outputs in order.
	Outputs []*StepOutput `protobuf:"bytes,3,rep,name=outputs" json:"outputs,omitempty"`
}

func (m *SubpipelinePipelineDescriptionStep) Reset()         { *m = SubpipelinePipelineDescriptionStep{} }
func (m *SubpipelinePipelineDescriptionStep) String() string { return proto.CompactTextString(m) }
func (*SubpipelinePipelineDescriptionStep) ProtoMessage()    {}
func (*SubpipelinePipelineDescriptionStep) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{11}
}

func (m *SubpipelinePipelineDescriptionStep) GetPipeline() *PipelineDescription {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *SubpipelinePipelineDescriptionStep) GetInputs() []*StepInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *SubpipelinePipelineDescriptionStep) GetOutputs() []*StepOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// Used to represent a pipeline template which can be used to generate full pipelines.
// A placeholder is replaced with a pipeline step to form a pipeline. See README.md
// for restrictions on the number of them, their position, allowed inputs and outputs,
// etc.
type PlaceholderPipelineDescriptionStep struct {
	// List of inputs which can be used as inputs to resulting sub-pipeline. Resulting
	// sub-pipeline does not have to use all the inputs, but it cannot use any other inputs.
	Inputs []*StepInput `protobuf:"bytes,1,rep,name=inputs" json:"inputs,omitempty"`
	// A list of outputs of the resulting sub-pipeline.
	Outputs []*StepOutput `protobuf:"bytes,2,rep,name=outputs" json:"outputs,omitempty"`
}

func (m *PlaceholderPipelineDescriptionStep) Reset()         { *m = PlaceholderPipelineDescriptionStep{} }
func (m *PlaceholderPipelineDescriptionStep) String() string { return proto.CompactTextString(m) }
func (*PlaceholderPipelineDescriptionStep) ProtoMessage()    {}
func (*PlaceholderPipelineDescriptionStep) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{12}
}

func (m *PlaceholderPipelineDescriptionStep) GetInputs() []*StepInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PlaceholderPipelineDescriptionStep) GetOutputs() []*StepOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type PipelineDescriptionStep struct {
	// Types that are valid to be assigned to Step:
	//	*PipelineDescriptionStep_Primitive
	//	*PipelineDescriptionStep_Pipeline
	//	*PipelineDescriptionStep_Placeholder
	Step isPipelineDescriptionStep_Step `protobuf_oneof:"step"`
}

func (m *PipelineDescriptionStep) Reset()                    { *m = PipelineDescriptionStep{} }
func (m *PipelineDescriptionStep) String() string            { return proto.CompactTextString(m) }
func (*PipelineDescriptionStep) ProtoMessage()               {}
func (*PipelineDescriptionStep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

type isPipelineDescriptionStep_Step interface {
	isPipelineDescriptionStep_Step()
}

type PipelineDescriptionStep_Primitive struct {
	Primitive *PrimitivePipelineDescriptionStep `protobuf:"bytes,1,opt,name=primitive,oneof"`
}
type PipelineDescriptionStep_Pipeline struct {
	Pipeline *SubpipelinePipelineDescriptionStep `protobuf:"bytes,2,opt,name=pipeline,oneof"`
}
type PipelineDescriptionStep_Placeholder struct {
	Placeholder *PlaceholderPipelineDescriptionStep `protobuf:"bytes,3,opt,name=placeholder,oneof"`
}

func (*PipelineDescriptionStep_Primitive) isPipelineDescriptionStep_Step()   {}
func (*PipelineDescriptionStep_Pipeline) isPipelineDescriptionStep_Step()    {}
func (*PipelineDescriptionStep_Placeholder) isPipelineDescriptionStep_Step() {}

func (m *PipelineDescriptionStep) GetStep() isPipelineDescriptionStep_Step {
	if m != nil {
		return m.Step
	}
	return nil
}

func (m *PipelineDescriptionStep) GetPrimitive() *PrimitivePipelineDescriptionStep {
	if x, ok := m.GetStep().(*PipelineDescriptionStep_Primitive); ok {
		return x.Primitive
	}
	return nil
}

func (m *PipelineDescriptionStep) GetPipeline() *SubpipelinePipelineDescriptionStep {
	if x, ok := m.GetStep().(*PipelineDescriptionStep_Pipeline); ok {
		return x.Pipeline
	}
	return nil
}

func (m *PipelineDescriptionStep) GetPlaceholder() *PlaceholderPipelineDescriptionStep {
	if x, ok := m.GetStep().(*PipelineDescriptionStep_Placeholder); ok {
		return x.Placeholder
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PipelineDescriptionStep) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PipelineDescriptionStep_OneofMarshaler, _PipelineDescriptionStep_OneofUnmarshaler, _PipelineDescriptionStep_OneofSizer, []interface{}{
		(*PipelineDescriptionStep_Primitive)(nil),
		(*PipelineDescriptionStep_Pipeline)(nil),
		(*PipelineDescriptionStep_Placeholder)(nil),
	}
}

func _PipelineDescriptionStep_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PipelineDescriptionStep)
	// step
	switch x := m.Step.(type) {
	case *PipelineDescriptionStep_Primitive:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Primitive); err != nil {
			return err
		}
	case *PipelineDescriptionStep_Pipeline:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pipeline); err != nil {
			return err
		}
	case *PipelineDescriptionStep_Placeholder:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Placeholder); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PipelineDescriptionStep.Step has unexpected type %T", x)
	}
	return nil
}

func _PipelineDescriptionStep_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PipelineDescriptionStep)
	switch tag {
	case 1: // step.primitive
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrimitivePipelineDescriptionStep)
		err := b.DecodeMessage(msg)
		m.Step = &PipelineDescriptionStep_Primitive{msg}
		return true, err
	case 2: // step.pipeline
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SubpipelinePipelineDescriptionStep)
		err := b.DecodeMessage(msg)
		m.Step = &PipelineDescriptionStep_Pipeline{msg}
		return true, err
	case 3: // step.placeholder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PlaceholderPipelineDescriptionStep)
		err := b.DecodeMessage(msg)
		m.Step = &PipelineDescriptionStep_Placeholder{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PipelineDescriptionStep_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PipelineDescriptionStep)
	// step
	switch x := m.Step.(type) {
	case *PipelineDescriptionStep_Primitive:
		s := proto.Size(x.Primitive)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PipelineDescriptionStep_Pipeline:
		s := proto.Size(x.Pipeline)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PipelineDescriptionStep_Placeholder:
		s := proto.Size(x.Placeholder)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Pipeline description matches the D3M pipeline description.
// It serves two purposes: describing found pipelines by TA2 to TA3, and communicating pipeline
// templates by TA3 to TA2. Because of this some fields are reasonable only in one of those uses.
// They are marked with "TA2" or "TA3" in the comment, for fields to be set only by TA2 or only
// by TA3, respectivelly.
type PipelineDescription struct {
	// TA2: UUID of the pipeline. Templates do not have IDs. Matches "pipeline_id" from
	// "ListPipelinesResponse" and other related messages.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// "schema" field is not needed because it is fixed by the TA2-TA3 protocol version.
	// System which generated a pipeline or a template. Optional.
	Source *PipelineSource `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	// TA2: Timestamp when created. Templates do not have this timestamp.
	Created *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=created" json:"created,omitempty"`
	// Human friendly name of the pipeline. For templates it can be a hint to
	// TA2 how to name found pipelines. Optional.
	Name string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	// Human friendly description of the pipeline. Optional.
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	// List of users associated with the creation of the template and consequently of the pipeline.
	// TA2 can store this information into metalearning database. TA2 is not really expected to use
	// this information during pipeline search. TA2 should not have to understand TA3 users, mapping
	// between users and pipeline search IDs is something TA3 should handle. Optional.
	Users []*PipelineDescriptionUser `protobuf:"bytes,6,rep,name=users" json:"users,omitempty"`
	// In most cases inputs are datasets. But if TA3 wants to jut run a primitive, it can send a
	// template with only that primitive in the template, and then pass anything to its inputs during
	// execution. Here, we are describing possible inputs to the pipeline or template. Order matters.
	Inputs []*PipelineDescriptionInput `protobuf:"bytes,7,rep,name=inputs" json:"inputs,omitempty"`
	// Available outputs of the pipeline or template.
	Outputs []*PipelineDescriptionOutput `protobuf:"bytes,8,rep,name=outputs" json:"outputs,omitempty"`
	// Steps defining the pipeline.
	Steps []*PipelineDescriptionStep `protobuf:"bytes,9,rep,name=steps" json:"steps,omitempty"`
}

func (m *PipelineDescription) Reset()                    { *m = PipelineDescription{} }
func (m *PipelineDescription) String() string            { return proto.CompactTextString(m) }
func (*PipelineDescription) ProtoMessage()               {}
func (*PipelineDescription) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *PipelineDescription) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PipelineDescription) GetSource() *PipelineSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *PipelineDescription) GetCreated() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *PipelineDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PipelineDescription) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PipelineDescription) GetUsers() []*PipelineDescriptionUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *PipelineDescription) GetInputs() []*PipelineDescriptionInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PipelineDescription) GetOutputs() []*PipelineDescriptionOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *PipelineDescription) GetSteps() []*PipelineDescriptionStep {
	if m != nil {
		return m.Steps
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerArgument)(nil), "ContainerArgument")
	proto.RegisterType((*PrimitiveArgument)(nil), "PrimitiveArgument")
	proto.RegisterType((*DataArgument)(nil), "DataArgument")
	proto.RegisterType((*PrimitiveStepArgument)(nil), "PrimitiveStepArgument")
	proto.RegisterType((*StepInput)(nil), "StepInput")
	proto.RegisterType((*StepOutput)(nil), "StepOutput")
	proto.RegisterType((*PipelineSource)(nil), "PipelineSource")
	proto.RegisterType((*PipelineDescriptionUser)(nil), "PipelineDescriptionUser")
	proto.RegisterType((*PipelineDescriptionInput)(nil), "PipelineDescriptionInput")
	proto.RegisterType((*PipelineDescriptionOutput)(nil), "PipelineDescriptionOutput")
	proto.RegisterType((*PrimitivePipelineDescriptionStep)(nil), "PrimitivePipelineDescriptionStep")
	proto.RegisterType((*SubpipelinePipelineDescriptionStep)(nil), "SubpipelinePipelineDescriptionStep")
	proto.RegisterType((*PlaceholderPipelineDescriptionStep)(nil), "PlaceholderPipelineDescriptionStep")
	proto.RegisterType((*PipelineDescriptionStep)(nil), "PipelineDescriptionStep")
	proto.RegisterType((*PipelineDescription)(nil), "PipelineDescription")
}

func init() { proto.RegisterFile("pipeline.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcb, 0x6e, 0xd3, 0x4a,
	0x18, 0x8e, 0x9d, 0x4b, 0xeb, 0xdf, 0xe7, 0xa4, 0xed, 0x9c, 0x03, 0xb8, 0x51, 0xa5, 0x86, 0x29,
	0xa8, 0x5d, 0xa0, 0x69, 0x1b, 0xba, 0x40, 0xec, 0x7a, 0x01, 0x82, 0x90, 0xa0, 0x72, 0x0b, 0x0b,
	0x76, 0x6e, 0x3c, 0xb4, 0x16, 0x89, 0x6d, 0x8d, 0xc7, 0x95, 0xba, 0xe6, 0x49, 0x78, 0x03, 0xde,
	0x8a, 0x2d, 0x8f, 0x80, 0x3c, 0x9e, 0x8b, 0x93, 0x38, 0x4d, 0x77, 0xb6, 0xe7, 0xfb, 0xbf, 0xff,
	0xf2, 0xfd, 0xf3, 0x19, 0xba, 0x69, 0x94, 0xd2, 0x71, 0x14, 0x53, 0x92, 0xb2, 0x84, 0x27, 0xbd,
	0xfe, 0x75, 0x92, 0x5c, 0x8f, 0xe9, 0xbe, 0x78, 0xbb, 0xca, 0xbf, 0xed, 0x87, 0x34, 0x1b, 0xb1,
	0x28, 0xe5, 0x09, 0x93, 0x88, 0xed, 0x59, 0x04, 0x8f, 0x26, 0x34, 0xe3, 0xc1, 0x24, 0x95, 0x80,
	0xb5, 0x94, 0x45, 0x93, 0x88, 0x47, 0xb7, 0x8a, 0xd3, 0xbd, 0x0d, 0xc6, 0xb9, 0x7c, 0xc1, 0xbb,
	0xb0, 0x71, 0x9a, 0xc4, 0x3c, 0x88, 0x62, 0xca, 0x8e, 0xd9, 0x75, 0x3e, 0xa1, 0x31, 0x47, 0x08,
	0x5a, 0x61, 0xc0, 0x03, 0xcf, 0xea, 0x5b, 0x7b, 0x8e, 0x2f, 0x9e, 0xf1, 0x21, 0x6c, 0x9c, 0x2b,
	0x22, 0x0d, 0xdc, 0x02, 0x47, 0xb3, 0x4b, 0xb4, 0xf9, 0x80, 0x31, 0xfc, 0x73, 0x16, 0xf0, 0xe0,
	0x5e, 0xda, 0x5f, 0x16, 0x3c, 0xd2, 0xbc, 0x17, 0x9c, 0xa6, 0x1a, 0x3d, 0x00, 0x67, 0xa4, 0x2a,
	0x13, 0x21, 0xee, 0x00, 0x91, 0xb9, 0x5a, 0x87, 0x0d, 0xdf, 0xc0, 0x8a, 0x18, 0x53, 0x8f, 0x2d,
	0x63, 0xe6, 0xca, 0x2e, 0x62, 0x34, 0x0c, 0xed, 0xc8, 0xaa, 0x9a, 0x02, 0xfe, 0x2f, 0xa9, 0x96,
	0x3c, 0x6c, 0x94, 0x65, 0x9e, 0x00, 0xac, 0x06, 0xf2, 0x1b, 0xde, 0x06, 0xa7, 0x28, 0xf4, 0x7d,
	0x9c, 0xe6, 0xf5, 0x3d, 0x6d, 0x01, 0x14, 0x80, 0x4f, 0x39, 0x2f, 0x10, 0x5d, 0xb0, 0xa3, 0x50,
	0x9e, 0xdb, 0x51, 0x88, 0x9f, 0x41, 0xf7, 0x5c, 0x8a, 0x7c, 0x91, 0xe4, 0x6c, 0x44, 0x0b, 0x8e,
	0x38, 0x98, 0xa8, 0x01, 0x8a, 0x67, 0xfc, 0x01, 0x9e, 0x28, 0xd4, 0x99, 0x94, 0x3c, 0x4a, 0xe2,
	0xcf, 0x19, 0x65, 0xb3, 0x84, 0xa8, 0x0f, 0x6e, 0x68, 0x20, 0xa2, 0x6d, 0xc7, 0xaf, 0x7e, 0xc2,
	0x04, 0xbc, 0x1a, 0x32, 0xdd, 0xc0, 0x5c, 0xf2, 0x53, 0xd8, 0xac, 0xc1, 0xcb, 0x7e, 0x6a, 0x02,
	0xf4, 0x14, 0xec, 0xca, 0x14, 0x7e, 0x37, 0xa1, 0xaf, 0x47, 0x5f, 0x43, 0x57, 0x8c, 0x08, 0xed,
	0xcd, 0x2e, 0x90, 0x3b, 0x00, 0x23, 0x58, 0x55, 0xa6, 0x8f, 0xe0, 0x28, 0x05, 0x32, 0xcf, 0xee,
	0x37, 0xf7, 0xdc, 0xc1, 0x01, 0x59, 0xc6, 0x4f, 0x94, 0x90, 0xd9, 0x9b, 0x98, 0xb3, 0x3b, 0xdf,
	0x50, 0xa0, 0xe7, 0xb0, 0x92, 0x88, 0x86, 0x32, 0xaf, 0x29, 0xd8, 0x5c, 0x62, 0x44, 0xf3, 0xd5,
	0x19, 0xba, 0x04, 0xf7, 0xe6, 0x2e, 0xa5, 0x2c, 0x0d, 0x58, 0x30, 0xc9, 0xbc, 0x96, 0x80, 0x0e,
	0x96, 0x27, 0x1e, 0x9a, 0xa0, 0x32, 0x75, 0x95, 0x06, 0x11, 0x68, 0xe7, 0x19, 0x65, 0x99, 0xd7,
	0x16, 0x7c, 0x1e, 0x59, 0xa0, 0xb5, 0x5f, 0xc2, 0x7a, 0x97, 0xd0, 0x9d, 0xee, 0x04, 0xad, 0x43,
	0xf3, 0x3b, 0xbd, 0x93, 0x22, 0x14, 0x8f, 0xe8, 0x05, 0xb4, 0xc5, 0xc5, 0x96, 0x7b, 0xff, 0x98,
	0xd4, 0x5e, 0x2b, 0xbf, 0x04, 0xbd, 0xb6, 0x5f, 0x59, 0xbd, 0xb7, 0xb0, 0x3e, 0x5b, 0x66, 0x0d,
	0xef, 0xd6, 0x34, 0x6f, 0x87, 0x7c, 0x29, 0xde, 0x2a, 0x3c, 0xf8, 0xa7, 0x05, 0xf8, 0x22, 0xbf,
	0x52, 0xd6, 0xb5, 0x48, 0xeb, 0x03, 0x58, 0x55, 0x10, 0x29, 0xf5, 0xff, 0x75, 0x7d, 0xfb, 0x1a,
	0x85, 0x30, 0x74, 0xa2, 0x58, 0x48, 0x54, 0x0a, 0x0e, 0x44, 0x5f, 0x3c, 0x5f, 0x9e, 0x3c, 0x50,
	0x47, 0x9c, 0x00, 0x3e, 0x1f, 0x07, 0x23, 0x7a, 0x93, 0x8c, 0x43, 0xca, 0x16, 0x95, 0x68, 0x12,
	0x5a, 0x0f, 0x49, 0x68, 0xdf, 0x93, 0xf0, 0x8f, 0x55, 0x7b, 0x83, 0x45, 0x9a, 0xe3, 0xf9, 0xad,
	0x7f, 0xba, 0x74, 0xa5, 0xa6, 0x5d, 0xeb, 0xb8, 0x32, 0xcc, 0x52, 0x98, 0x1d, 0xb2, 0x5c, 0x83,
	0x61, 0xa3, 0x32, 0xdd, 0x77, 0xe0, 0xa6, 0x66, 0x24, 0xd2, 0xff, 0x76, 0xc8, 0xf2, 0x31, 0x0d,
	0x1b, 0x7e, 0x35, 0xf2, 0xa4, 0x03, 0xad, 0x8c, 0xd3, 0x14, 0xff, 0x68, 0xc2, 0x7f, 0x35, 0x21,
	0x73, 0x86, 0xb5, 0x0b, 0x9d, 0x4c, 0x38, 0x9f, 0xac, 0x7c, 0x8d, 0x4c, 0x1b, 0xa2, 0x2f, 0x8f,
	0xd1, 0x11, 0xac, 0x8c, 0x18, 0x0d, 0x38, 0x0d, 0x65, 0x75, 0x3d, 0x52, 0xfe, 0xed, 0x88, 0xfa,
	0xdb, 0x91, 0x4b, 0xf5, 0xb7, 0xf3, 0x15, 0x54, 0x1b, 0x54, 0xab, 0x62, 0x50, 0x33, 0x1e, 0xd9,
	0x9e, 0xf3, 0x48, 0x73, 0x25, 0x3b, 0x0f, 0xba, 0x92, 0xe8, 0x50, 0xaf, 0xca, 0x8a, 0x08, 0xd8,
	0x24, 0x8b, 0x2c, 0x56, 0x6f, 0xce, 0x91, 0xd9, 0x9c, 0x55, 0x11, 0xd3, 0x23, 0x0b, 0x6d, 0xd6,
	0x38, 0x10, 0x81, 0x76, 0x31, 0xdd, 0xcc, 0x73, 0x16, 0x17, 0x56, 0xa8, 0xe2, 0x97, 0xb0, 0x13,
	0xf8, 0xaa, 0x25, 0xbe, 0xea, 0x88, 0x39, 0xbd, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x39, 0x9e,
	0x2f, 0xeb, 0x57, 0x08, 0x00, 0x00,
}