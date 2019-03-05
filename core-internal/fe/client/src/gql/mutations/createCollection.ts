import gql from 'graphql-tag';

const CREATE_COLLECTION = gql`
  mutation {
    createCollection(
      input: {
        name: "asdas"
        host: "qwertr"
        fieldValues: {
          valueType: INT32
          fieldName: "ertgdf"
          isUnique: true
          isIndexed: true
        }
      }
    ) {
      id
    }
  }
`;
export default CREATE_COLLECTION;
