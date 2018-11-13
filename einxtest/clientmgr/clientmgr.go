package clientmgr

import (
	"github.com/Cyinx/einx"
	"github.com/Cyinx/einx/slog"
)

type Agent = einx.Agent
type AgentID = einx.AgentID
type EventType = einx.EventType
type Component = einx.Component
type ComponentID = einx.ComponentID

type ClientMgr struct {
	client_map map[AgentID]Agent
	tcp_link   Component
}

var Instance = &ClientMgr{
	client_map: make(map[AgentID]Agent),
}

func (this *ClientMgr) GetClient(agent_id AgentID) (Agent, bool) {
	client, ok := this.client_map[agent_id]
	return client, ok
}

func (this *ClientMgr) OnLinkerConneted(id AgentID, agent Agent) {
	this.client_map[id] = agent //新连接连入服务器
}

func (this *ClientMgr) OnLinkerClosed(id AgentID, agent Agent) {
	delete(this.client_map, id) //连接断开
}

func (this *ClientMgr) OnComponentError(c Component, err error) {

}

func (this *ClientMgr) OnComponentCreate(id ComponentID, component Component) {
	this.tcp_link = component
	component.Start()
	slog.LogInfo("tcp", "Tcp sever start success")
}
