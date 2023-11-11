// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

contract Highway {
    // EvidenceEntity 存证结构体
    struct EvidenceEntity {
        // Id 业务流水号
        string Id;
        // Hash 哈希值
        string Hash;
        // TxId 存证时交易ID
        string TxId;
        // BlockHeight 存证时区块高度
        uint BlockHeight;
        // Timestamp 存证时区块时间
        string Timestamp;
        // Metadata 可选，其他信息；具体参考下方 Metadata 对象。
        MetadataEntity Metadata;
    }

    // MetadataEntity 可选信息建议字段，若包含以下相关信息存证，请采用以下字段
    struct MetadataEntity {
        // HashType 哈希的类型，文字、文件、视频、音频等
        string HashType;
        // HashAlgorithm 哈希算法，sha256、sm3等
        string HashAlgorithm;
        // Username 存证人，用于标注存证的身份
        string Username;
        // Timestamp 可信存证时间
        string Timestamp;
        // ProveTimestamp 可信存证时间证明
        string ProveTimestamp;
        // 存证内容
        string Content;
        // 其他自定义扩展字段
        // ...
    }

    // Evidence 存证
    // @param id 必填，流水号
    // @param hash 必填，上链哈希值
    // @param metadata 可选，其他信息；比如：哈希的类型（文字，文件）、文字描述的json格式字符串，具体参考下方 Metadata 对象。
    // @return error 返回错误信息
    function Evidence(
        string memory _id,
        string memory _hash,
        MetadataEntity memory _metadata
    ) public {}

    // EvidenceBatch 批量存证
    // @param evidences 必填，存证信息
    function EvidenceBatch(EvidenceEntity[] memory evidences) public {}

    // UpdateEvidence 根据ID更新存证哈希和metadata
    // @param id 必填，已经在链上存证的流水号。 如果是新流水号返回错误信息不存在
    // @param hash 必填，上链哈希值。必须与链上已经存储的hash不同
    // @param metadata 可选，其他信息；具体参考下方 Metadata 对象。
    // @desc 该方法由长安链社区志愿者@sunhuiyuan提供建议，感谢参与
    function UpdateEvidence(
        string memory _id,
        string memory _hash,
        MetadataEntity memory _metadata
    ) public {}

    // ExistsOfHash 哈希是否存在
    // @param hash 必填，上链的哈希值
    // @return exist 存在：true，"true"；不存在：false，"false"
    // @return err 错误信息
    function ExistsOfHash(
        string memory _hash
    ) public view returns (bool exist, string memory error) {}

    // ExistsOfId 哈希是否存在
    // @param id 必填，上链的ID值
    // @return exist 存在：true，"true"；不存在：false，"false"
    // @return err 错误信息
    function ExistsOfId(
        string memory _id
    ) public view returns (bool exist, string memory error) {}

    // FindByHash 根据哈希查找
    // @param hash 必填，上链哈希值
    // @return evidence 上链时传入的evidence信息
    // @return err 返回错误信息
    function FindByHash(
        string memory _hash
    )
        public
        view
        returns (EvidenceEntity memory evidence, string memory error)
    {}

    // FindById 根据id查找
    // @param id 必填，流水号
    // @return evidence 上链时传入的evidence信息
    // @return err 返回错误信息

    function FindById(
        string memory _id
    )
        public
        view
        returns (EvidenceEntity memory evidence, string memory error)
    {}

    // FindHisById 根据ID流水号查找存证历史(可以使用go合约接口：sdk.Instance.NewHistoryKvIterForKey或NewIterator实现)
    // @param id 必填，流水号
    // @return string 上链时传入的evidence信息的各个版本JSON数组对象。如果之前上链没有调用过updateEvidence、效果等同于findById，数组大小为1；
    //                如果之前上链调用过updateEvidence，则结果数组长度>1。
    // @return error 返回错误信息
    // @desc 该方法由长安链社区志愿者@sunhuiyuan提供建议，感谢参与
    function FindHisById(
        string memory _id
    ) public returns (EvidenceEntity memory evidence, string memory err) {}
}
