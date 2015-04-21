create table IF NOT EXISTS resource_tab(
  res_id  varchar(50) not null COMMENT '资源ID',
  res_name varchar(100) not null COMMENT '资源名称',
  owner_acid bigint(20) not null COMMENT '拥有者账号ID',
  operator_acid bigint(20) not null COMMENT '授权者',
  status   integer not null COMMENT '资源状态',
  create_time integer not null COMMENT '创建时间',
  PRIMARY KEY (`res_id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='注册中心-资源表' ;

create table IF NOT EXISTS res_privilige_tab(
  res_id  varchar(50) not null COMMENT '资源ID',
  pri_name varchar(100) not null COMMENT '权限名称',
  pri_desc varchar(100) not null COMMENT '权限描述',
  PRIMARY KEY (`res_id`,`pri_name`)
)ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='资源权限对应表' ;

create table IF NOT EXISTS account_tab(
  acid bigint(20) not null AUTO_INCREMENT COMMENT '账号ID',
  ac_name varchar(100) not null COMMENT '账户名称',
  ac_password varchar(50) not null COMMENT '账户密码',
  email varchar(50) not null COMMENT '电子邮件',
  mobile varchar(50) not null COMMENT '手机号码',
  status   integer not null COMMENT '账户状态',
  create_time   integer not null,
  PRIMARY KEY (`acid`)
)ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='认证中心-账号表'  AUTO_INCREMENT=1  ;

create table IF NOT EXISTS openid_tab(
  res_id  varchar(50) not null COMMENT '资源ID',
  acid bigint(20) not null COMMENT '账号ID',
  openid   varchar(50) not null,
  PRIMARY KEY (`res_id`,`acid`)
)ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='账号OpenID对应表' ;

create table IF NOT EXISTS authorization_tab(
  res_id  varchar(50) not null COMMENT '资源ID',
  acid bigint(20) not null COMMENT '账号ID',
  priviliges varchar(512) not null COMMENT '权限名称',
  operator_acid varchar(50) not null COMMENT '权限名称',
  PRIMARY KEY (`res_id`,`acid`)
)ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='授权中心-授权关系表' ;

create table IF NOT EXISTS authorization_tab(
  res_id  varchar(50) not null COMMENT '资源ID',
  app_id bigint(20) not null COMMENT '资源APP应用的id',
  app_key varchar(50) not null COMMENT '资源APP客户的id',
  PRIMARY KEY (`res_id`)
)ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='授权中心-授权关系表' ;


